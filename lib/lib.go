// Package lib implements core qri business logic. It exports
// canonical methods that a qri instance can perform regardless of
// client interface. API's of any sort must use lib methods
package lib

import (
	"encoding/gob"
	"fmt"
	"net"
	"net/rpc"
	"os"
	"path/filepath"

	golog "github.com/ipfs/go-log"
	"github.com/mitchellh/go-homedir"
	"github.com/qri-io/ioes"
	"github.com/qri-io/qfs"
	ipfs "github.com/qri-io/qfs/cafs/ipfs"
	"github.com/qri-io/qfs/httpfs"
	"github.com/qri-io/qfs/localfs"
	"github.com/qri-io/qfs/muxfs"
	"github.com/qri-io/qri/config"
	"github.com/qri-io/qri/p2p"
	"github.com/qri-io/qri/repo/fs"
	"github.com/qri-io/qri/repo/profile"
	"github.com/qri-io/registry/regclient"
)

var log = golog.Logger("lib")

// VersionNumber is the current version qri
const VersionNumber = "0.7.1-dev"

// Requests defines a set of library methods
type Requests interface {
	// CoreRequestsName confirms participation in the CoreRequests interface while
	// also giving a human readable string for logging purposes
	CoreRequestsName() string
}

func init() {
	// Fields like dataset.Structure.Schema contain data of arbitrary types,
	// registering with the gob package prevents errors when sending them
	// over net/rpc calls.
	gob.Register([]interface{}{})
	gob.Register(map[string]interface{}{})
}

// Receivers returns a slice of CoreRequests that defines the full local
// API of lib methods
func Receivers(node *p2p.QriNode) []Requests {
	return []Requests{
		NewDatasetRequests(node, nil),
		NewRegistryRequests(node, nil),
		NewLogRequests(node, nil),
		NewExportRequests(node, nil),
		NewPeerRequests(node, nil),
		NewProfileRequests(node, nil),
		NewSearchRequests(node, nil),
		NewRenderRequests(node.Repo, nil),
		NewSelectionRequests(node.Repo, nil),
	}
}

// QriOptions provides details to New. New will alter QriOptions by applying
// any provided Option functions
// to distinguish "Options" from "Config":
// * Options contains state that can only be determined at runtime
// * Config consists only of static values stored in a configuration file
// Options may override config in specific cases to avoid undefined state
type QriOptions struct {
	Cfg      *config.Config
	Streams  ioes.IOStreams
	QriPath  string
	IPFSPath string
}

// Option is a function that manipulates config details when fed to New()
type Option func(o *QriOptions) error

// OptDefaultQriPath configures the directory to read Qri from, defaulting to
// "$HOME/.qri", unless the environment variable QRI_PATH is set
func OptDefaultQriPath() Option {
	return func(o *QriOptions) error {
		path := os.Getenv("QRI_PATH")
		if path == "" {
			dir, err := homedir.Dir()
			if err != nil {
				return err
			}
			path = filepath.Join(dir, ".qri")
		}
		o.QriPath = path
		return nil
	}
}

// OptDefaultIPFSPath configures the directory to read IPFS from, defaulting to
// "$HOME/.ipfs", unless the environment variable IPFS_PATH is set
func OptDefaultIPFSPath() Option {
	return func(o *QriOptions) error {
		path := os.Getenv("IPFS_PATH")
		if path == "" {
			dir, err := homedir.Dir()
			if err != nil {
				return err
			}
			path = filepath.Join(dir, ".ipfs")
		}
		o.IPFSPath = path
		return nil
	}
}

// OptLoadConfigFile loads a configuration from a given path
func OptLoadConfigFile(path string) Option {
	return func(o *QriOptions) error {
		// default to checking
		if path == "" && o.QriPath != "" {
			path = filepath.Join(o.QriPath, "config.yaml")
		}

		if err := LoadConfig(o.Streams, path); err != nil {
			return err
		}
		ConfigFilepath = path
		return nil
	}
}

// OptIOStreams sets the input IOStreams
func OptIOStreams(streams ioes.IOStreams) Option {
	return func(o *QriOptions) error {
		o.Streams = streams
		return nil
	}
}

// OptStdIOStreams sets treams to std, stdout, & stderr
func OptStdIOStreams() Option {
	return func(o *QriOptions) error {
		o.Streams = ioes.NewStdIOStreams()
		return nil
	}
}

// New creates a new Qri handle, if no Option funcs are provided, New uses
// a default set of Option funcs
func New(opts ...Option) (qri *Qri, err error) {
	o := &QriOptions{}
	if len(opts) == 0 {
		// default to a standard composition of Option funcs
		opts = []Option{
			OptStdIOStreams(),
			OptDefaultQriPath(),
			OptDefaultIPFSPath(),
			OptLoadConfigFile(""),
		}
	}
	for _, opt := range opts {
		if err = opt(cfg); err != nil {
			return
		}
	}

	cfg := o.Cfg
	if cfg == nil {
		err = fmt.Errorf("no configuration details provided")
		return
	}
	if err = cfg.Validate(); err != nil {
		return
	}

	if cfg.RPC.Enabled {
		addr := fmt.Sprintf(":%d", cfg.RPC.Port)
		if conn, err := net.Dial("tcp", addr); err != nil {
			err = nil
		} else {
			qri = &Qri{
				cfg: cfg,
				rpc: rpc.NewClient(conn),
			}
			return qri, err
		}
	}

	var store *ipfs.Filestore
	fsOpts := []ipfs.Option{
		func(c *ipfs.StoreCfg) {
			c.FsRepoPath = cfg.Store.Path
			// c.Online = online
		},
		ipfs.OptsFromMap(cfg.Store.Options),
	}
	if store, err = ipfs.NewFilestore(fsOpts...); err != nil {
		return
	}

	var pro *profile.Profile
	if pro, err = profile.NewProfile(cfg.Profile); err != nil {
		return
	}

	var rc *regclient.Client
	if cfg.Registry != nil && cfg.Registry.Location != "" {
		rc = regclient.NewClient(&regclient.Config{
			Location: cfg.Registry.Location,
		})
	}

	fsys := muxfs.NewMux(map[string]qfs.PathResolver{
		"local": localfs.NewFS(),
		"http":  httpfs.NewFS(),
		"cafs":  store,
		"ipfs":  store,
	})

	repo, err := fsrepo.NewRepo(store, fsys, pro, rc, qriRepoPath)
	if err != nil {
		return
	}

	node, err := p2p.NewQriNode(repo, cfg.P2P)
	if err != nil {
		return
	}
	node.LocalStreams = o.Streams

	return &Qri{cfg: cfg, node: node}, nil
}

// Qri is a single handle for accessing all of Qri's subsystems
// create one with New
type Qri struct {
	cfg  *config.Config
	node *p2p.QriNode
	rpc  *rpc.Client
}

// Datasets provides methods for working with Datasets
func (q *Qri) Datasets() *DatasetRequests {
	return NewDatasetRequests(q.node, q.rpc)
}

// Registries provides methods for working with Registries
func (q *Qri) Registries() *RegistryRequests {
	return NewRegistryRequests(q.node, q.rpc)
}

// Logs provides methods for working with Logs
func (q *Qri) Logs() *LogRequests {
	return NewLogRequests(q.node, q.rpc)
}

// Exports provides methods for working with Exports
func (q *Qri) Exports() *ExportRequests {
	return NewExportRequests(q.node, q.rpc)
}

// Peers provides methods for working with Peers
func (q *Qri) Peers() *PeerRequests {
	return NewPeerRequests(q.node, q.rpc)
}

// Profiles provides methods for working with Profiles
func (q *Qri) Profiles() *ProfileRequests {
	return NewProfileRequests(q.node, q.rpc)
}

// Searchs provides methods for working with Searchs
func (q *Qri) Searchs() *SearchRequests {
	return NewSearchRequests(q.node, q.rpc)
}

// Renders provides methods for working with Renders
func (q *Qri) Renders() *RenderRequests {
	return NewRenderRequests(node.q.node, q.rpc)
}

// Selections provides methods for working with Selections
func (q *Qri) Selections() *SelectionRequests {
	return NewSelectionRequests(node.q.node, q.rpc)
}
