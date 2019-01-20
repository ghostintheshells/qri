// Package test contains predefined set of PeerInfo for testing.
// Generating private keys is computationally expensive, which slows down tests dramatically.
// Using a cache of pre-calculated private keys defined here to keeps tests fast.
package test

import (
	"encoding/base64"

	crypto "gx/ipfs/QmPvyPwuCgJ7pDmrKDxRtsScJgBaM5h4EpRL2qQJsmXf4n/go-libp2p-crypto"
	peer "gx/ipfs/QmTRhk7cgjUf2gfQ3p2M9KPECNZEW9XUrmHcFCgog4cPgB/go-libp2p-peer"
)

// EncodedPeerInfo holds encoded private keys and peerIDs for test peers.
type EncodedPeerInfo struct {
	Base64PrivKey string
	B58PeerID     string
}

// PeerInfo holds deencoded keys and peerIDs for test peers.
type PeerInfo struct {
	PrivKey        crypto.PrivKey
	PubKey         crypto.PubKey
	PeerID         peer.ID
	EncodedPrivKey string
	EncodedPeerID  string
}

var encoded = []EncodedPeerInfo{
	{"CAASqAkwggSkAgEAAoIBAQChp1HiZxTsLQCaHmW3/cc2ZDZpgLwn5o1/nZPgqT7SyXHP5bn7GQMG3kPEQWcl4nhtLX9hkrBEskHrdIlqp9zXFMwBfat+qfzCylGC/QBDF7wT9umLd7nbq7pAxQXteXgntt2Zhg4gE/kEk7vIyL+P9KpWJZ/yjpykgsDC7NPnrr8qZBo2tL0F4w+33nZhEx7Pp7Rnaq22JM8rF+NHCgSkUh63lp7Vhwm9PQoGtt0XTnEKxrMQnUme/IhGNxs84RphxHc5+nW6jYjgm5bcJonGyPU7bq+v51Mr2Ol4RT3L9ZNJgz0SWTSmAtiBLx2ryLrTjmDPSvN7wLm9sWEdWmRVAgMBAAECggEBAJMumrl+jWgz2TZ5sreBEp6NQ5VvpuDVY8PrnzaQIikdTMizK1BaB417VUwdGGM//dG5+R7HxkHl42sT4gH/8GzL/Krm1vwunXplZy3SWSi9NXsf9qgLTGebxasvOCRt0l6mesFLcxT12ma2c+VuEixp4aUqAKWB/1Ex03wm0RFBcSttPHe5ODW8Eaz+ZU8cpObEcZdCIPVxeWqLVdkAImOmsknL0EAxP8Wo/V6Rh5Cg4PnwnfJiQ45C+m6h7NTIw0H4UOncv7EBABra6LqF6Uoda9vmv8CpwaXwR557DPchQglFjtm48jWGeVKO3Zyutizu420eRrFZ0GmJo5flvkkCgYEA0SLysOZNxDgjYA0ihVYL6UbCvYUSADuDyTWREOUiRfmxAmS1xN9o7fieCJnA4aAAnSugtT2BI7HEqT1lLz0YF8NRDKL07TNbkmNLIHXBbXA5saf10N2juhflfIm5/b/W9lC3QsngMR27J25Ztqof6Ur36bIKJ6Y6XvYdlkkZkc8CgYEAxeCHUWMvtHtBID9ZOtrZRNhNJ/uz+2rzVSPd6ZdhEUWsvv/0p7JXmSAp2eoJDDKHeSnVxcxQMqhq0/edUSSzSvDpWha8UU4N8hRpu+M0XZNke0ijhpK6NIqNHPvZdsyFD0VR1Vaj2Ruy+pzih6PhqSnn2ZwvpQJAwBnqc2VCJJsCgYAkQr33hAbpxZ4EkmJw4elwye8L8x2a4rbH1TzQxBm8Lj3Nn26Qsve7gwbLkPULabWRirXzlrVkXfcuLNH1bc9Wl2vfGAYFdokjCYpGF4SxF+s47VlGnJc9tdT5UdvorjF0RaxwrRXtDi2b+Zsee8LKrU/sugzesQif3GZm30fKqwKBgQCQHwHP+HMFfAQqLZma8UzwBK7loUEsrHAAoff+K8CKKPoxvxD9lzqQD8oLqpbeaGsdh6fowe/jhaERM7dEI3vm6GK9t/N/MF+d4tpD+67nPPQhiv13haTTodo3swNnsHx1a+K3hLwf5DnOqLehXW59nET+zPAyudpZUEbft2+eYwKBgCMS6SitXwa2UjFNgkMAaOeJjkjnUKcr1tO/zPtaYPugKgkMQB890q4dcq5rnG2onhJ7hkoMwcrFugbD2nub9AIkaMc6Y46jyh2mSeA0337MpoMp99Jmp2/B1rouYo4IRS25b7jk22yjV8ARCzsxFVQxEwA1Lg8YpaXaifuI+/2O", "QmeL2mdVka1eahKENjehK6tBxkkpk5dNQ1qMcgWi7Hrb4B"},
	{"CAASqAkwggSkAgEAAoIBAQDE6jTAXSR6TbBb2rfkosYR0XIrmR9sH/0HJI75xq55DIJGBVcl0ki+9nKLUOCi/pC487BP4ZzSsdTctThrINbIYEu3xe1lggFwNvzFlSag8sc/F97nooGbXpXIBUngVIVqdUT8idNAPZYiO1fEnxcb51hGP9K7h1cnFAfwzsKqJhs/BF6EioS7l0uXlKU6BPoRLVcjGOYtSJAgHOkhlaa+ISeDHK5Iy8ggoasZ0lezHKW9A3PZ5vrIA+PwiJSRihTMm966kzyizDY3/hI5tdEfJYJMsdqyklyaSsq55RY62otsL6wLKjGctyPV30ZLUSdI2kZuyO2w/ok95GOMJqWPAgMBAAECggEBAKwX89pajNLGqubcE/MhvvE7lwg7XpbkrgJcFQh+d2UbZY9Eg5FuYl1ijWDsYiaRTHIXp3NoveH1wQ7S4mfd31hnsEUAGiWopREpPWiAna3z/+ZIOms+Pv9Gfqi81n/T3nXX317GJXXzXQ61xlL0pwGgAioDBW0XLzfb7cSrLr373N5BQp6j/Et8C6oWnT48LOBr7TIE3unjVC/g64LjRJ24Ry5XZAJno8kUjvX/qg6LB2sqLLbA6R8FeCwszRCXTBzIdTbadX4FkPKpI7NrfCfOLG/L+Zf2LJEfAw/vBi+F7jBq4Rt9qvyjJj7/1oJrJ6QGESmFdbPGX+HgbwF0a0ECgYEA+t7IJKQf8deZceMDJ/jucy1ZSJwW9iv9CYMt6rxCBFNaOpEk1E+zP6Kz5wsqPXUSf5dTpDBH/vm0D255rmwVz0Tpx7xRKokvAYcjTexZehxzhqPdZymcPRZIF9V7Iyln9qUAJuLUf2MwnMbMl69/Poz/dq+iK3+HpnLnxg1RBRcCgYEAyPD8l9h8C7Jp1uOPb65IVwbm7dZJ7e/My3vjsoVewcoHZA6P4N21MPllI61VRZr3Z1uMDhbwmUvp2wtOTtXuAImERs2PFszG9NQQyNvRqNMzMFH+MOcXiQ3/Ws9zfK8reKmufD5ZNQe+HcgWLfIuy7iOt0p7xggTYGewZVlYnkkCgYBPWkzAmlGoc+QLjB0hdbInKH8HYqg4se1WJvJNP8M6DwuJXwPhTFyMknCJcpSn3/I7/aftVYBQfLeh8fX3YCT97PRtw3mBFOeTeiWGrm4XHAzG1+peiWDsSbIAJ/zNQHmsIMENi85fhQaJcLCiglajeIIODrwjOjG0SsBZezjXfQKBgHWl+w5glsg2bpd9ZsbJsNsbVGveMizYYPymjbtBMSifQ9KGYCEVTffdnSTVYH6/a6kdRZQeREJM2x//r5qi0JWJ7mOSCPwda0N/QlCHu2pwNaFN8FjrhLEe++pMWd6fpQEhv+JIkuxkmyBOvQWrrVBjv1N7jZp1sfqY2wOL20HZAoGBANMrTbs2YZW+Jy3WmGEm/MMI5VKr1ajbyJoFAEOVEggSAchI3B/E9HhTN7vOC44WMHgOVgCdZfqQoRjfJIOqvcVzPsFXUV8hi0kCsk6+RWMuiIwW1LDGi/LXJT6gkkfzPs4gAyCJ3tEmPWkOlgtORhb1zLRJsuudeHst5Of12A/b", "QmWYgD49r9HnuXEppQEq1a7SUUryja4QNs9E6XCH2PayCD"},
	{"CAASqAkwggSkAgEAAoIBAQCfGmVPd7E3bItgeHOHYVN6aOtah4G6nvsOoJnu31nA6ZnPYAwop7oCcpBDHk49u9e6KCfiZzyvDY2Kkl2ioddqoOL0kIzUAZVqcuQ0k8xBmkaqFItktsYLWtBM0KeCKJS0fCoEFV9AJpgFEVEyJfPHXENpDvfDW3kLQT9Cxki/o6r7ttlAuyb0O7iFQk9Lw1yHS1gpCAiiXulNsQYFtTb7aIGHGtjo09bvB9JX/k46vINDYPOAApVX6vu075fwiAFdtYF5P+7oofYSTmNfCA4EdvE4ZUVdk9kmd1go8tXniHf/kpLCO9LoZhbaGk8bP/MrKESnEmv9IDpyHpm7ZFSVAgMBAAECggEAboXRltCYxN2cPJmv931+lexIggzb654A3PpntG5nVQFxnYMlLyDEsGqRGG72/x98Do9SlY4Ns7UAfCCM9yriD+pPfoDjs2qeNuI08Oky3Oek6wV2h7IIBV9Cj/nqQxqZon8WWNvX3SJyPCL5epKus/C9yz9gkc77WPwsS+CeLLEW8r/YLTojXNTv3+9lvKVs3L7YoQ8jzLRGrB3SvuAEdJO0K7pj7ETWByvSMo81SS+z4dVXjLnjV0GTXB17Lc6sYd2pZb2UjRqrdup5wih7CO7XRozw5YZWMWwKuuujtnJOuEZj59T/aMJ3ZbO+Ra28mByoOPSSLH85iDutvJgmSQKBgQDJTC9YTohY3Di4S0VoM0OnflEHEuctUsolPDdJK8QJv8wW42dkyw4u3yCHxkAw5XhSLH2IP9AKYiUtkdlHnabK51IvNM/UM01i7mJ3RWsXINmjA5nGsNIu2r76l/TS5wsd0EYRmIzWsTNArVcmhmO+QFDFfmEeYopwhL1YGlI7TwKBgQDKVtYcOrDCsiEumffVHNJ4uzjttAcqTk9SyChxxzlTWsMtrSUN/Bip9nQzMoMbI0g8Ax8m0Qfc1ow4eQbdg5mBUJ9hf4sGwOVeMemOhu/xvdmbamenMVUbJ2pteXleep/k+OqpGOYhVHCRGVhDK8knZm6LTqIaXvAhW/VKtvTo2wKBgE0uUNjjA5ROm9DNy4bwYPhPjnHP5aOTIr60H96GUVwLlg0QlAwpbv65vpdQhDw7P68CL/+VuAbDKs4WJeE3qUQS/YfsD+Ok5/8Ot+JTU4RCrC+8qbFtCqm8ZY1fxhc3TyPXP4Zvn1CSELyzub6juIqxzkZq9oHX9oy92Ht3XCAxAoGBAL7Z2WsXENMUIahLIjCp1vx52+CaBogpBTkYAs4PFEtDOQZEYscmGj12cXQ5GODoDqJMb321fK0Y0XoS7h7SmmAXI75bB09/ctm+SoZdUMW8RR0K6GoTZisOqpxID+rFzzuybeTBz49wjhzOwynYiwvW2j4cFVq6YdWVX3Dx0WAxAoGBAMY2aM5tXpTNh84+hq8UQKp5dZWK/KpwA4B+P4MLGacRUMyC9zojpmoiaiC1OOc2nA1TDrovr9wMEpKEPbnop7nZWnrJpgglU98zg/+bgzBZe5JAQ7Y0ysx4z7sUfOXFnmw09QCfaQaaczcSR5OtWnpfj9CTsTpAeJ6tRw/K9hfd", "QmPeFTNHcZDr3ZFEfFfepxS5PqHAmfBRGQNPJ389Cwh1as"},
	{"CAASqQkwggSlAgEAAoIBAQDhj23A0l8eI+BoKFu96QD/TPOcqUgu0072m0p7WoMVv+JvvTK/2NiU1gculdQ86a+mqhjX3DseU/4a1UcsiZLVeuGSa5SvJtke6YyHQAJ18v4dp7phUkaL/nwL/3qmVGk4yLmhKSj11Ty0179ISUTVWOXxRuWIbJrITjqRZCCHso5iJz+CsrV2vUZnaG/Emc17EnexU3WcZvza5psICzbgJbYSqdZAfpLNNqFKS38Uf6YSwnoWGyJli6qkgtpvBVQ0C3YNq/g5Jrkd30zkkdL6IndV5FYKgmqVqjeIv3coAP6+EltZKSPJUR83jSovFe66/psQSDpHIq1L7sUOOV9rAgMBAAECggEAOKP0sovhOMdXjbA8wWamSnW/YuKbbq6/UGvUez9SVX7W1KpU+Rnx4QYRzZ/wKM+hvR9Pl70E12vODlzqP4PYbGUsCcGcF5OQyBNrp3bqZ0X6i+XXhqLLdlIUQksIOIZP2eOFwbN9RqYiVWMr9HH5p/1n2HuuHeDvyUbjrCjRhmFWcQxu/R9ESEaiw7Cbtb8WI/3yyR9cq//gHX1AlEn/CG5vQlrI9vB/URRmI86i9cS5LioHyex8VY01ZiKdSve4OyxSJA2k5fhkq9wT3QyHaH57/QfLYqm4UfNb9lgqIaUqcrtzwAQEfYm51f+23I35Rno1O114UyZL2SC6topnkQKBgQDyqnEq0s9wTNlud8VwAPv/pRk+YnPSSvOk8iLbwfDdoF06ki106xMw/W2sMOyXLlQCqwLeojezqnpVwiIhOQIaGlWbbuQXndBEi1OGiFdUDu3mru1vKW9iw+tDuR7lkSTCNSjVuPva0eGCscaA/Wx+BTzoUgdjkd2p3OUMgPi9EwKBgQDt9F1lsUSsceoEuz1GMdzeWtFXr/BVGTlJq4wad+s5tgu4DETDfo6cmmr3beNltQzuomPgRZMpQikSAUlqFe/DgWlcnQLBpDMuPcF+MgkNfH9Fgp07frrIizyoxmaV0sqwLWpI06d/ivZ5MyqvqTW9XIWMg5D4jDFh1apZ+A5XSQKBgQCBJcLq1p0+GDhT+XxjmrKDIRvpPr6Deg5nh9KTgIkveoyDgo6cvgtNtGLtFHCmGSru7JKvaEga360R/Srq4YtD1yYefgr4oq3X/Qqr+932R6fHcOu+kJ/OQZTxSxGtfezBS0d7T2MC11aclv6dYcKe4yzCO/3guR3urDheyjkXbwKBgQC3EwlAkjBZ1JLN7Rjphrd9w7XIly566psu2PND8ftiPXOquJW3Kwmh5xRhJYSM3c0DVKEvDZh5Z9OdbKwFGLeNXWWowYw7W4+dBhp2cxyP6bcpzaIXAhvG64lR+MM4hiM3hNl/CoiWEl4rRXiUCcW02RhO0XKaJ5JeyTr1WWn/MQKBgQC9MqCeRII/wGMi/SR8XJSjX47H3b4XYogBvWjYUmMnHzrZeQIJmgtJqCX8SogwOufKgepO2hz30H6+bY1fInXdavyB02+p7iiNhjTlryxozvpzF0rPb8vUMiCQ7LoegERvfCRBKtyYysnpy/TN82/AjTItWOTZW0yTCyvl5Z50YQ==", "QmbHm5S9hRsvZWE56XWA7dksfXDNABRwK6LaH6KS1z7uEw"},
	{"CAASqQkwggSlAgEAAoIBAQDT1VRrpQZNgiTXRRkCcbOM3WkIwIIdX2ZGN7Fg8IJMV1lG/bWTxQ/fgu48o6inDDYIr9dAnjRQa2tEKfhcSxJMfopKzaFYNW4V3NrPwU1TrFejMRd5khZk3GLsL+EkWhz6OoX6j98/wHra6e0YB5XvT5nUL7BdXHNsLt6vmRiumuQFS6Gkto7tuNGp72Vw9XJNxul8FPPv6BVBbhdFFVBQkXW9Txop3SuYK3OiXLucmODTOFgy34NXiPvTtdpAMnLxphDFZzZGZtfRIzLr4PH6aayiEZ9DOKx60Y4P3OvAYg8eBAcpfjcKRbzMZINQjzg+JyVd5wEpBulSNP1da9s9AgMBAAECggEBALI0bDGmgZfg3VdP5Ms+ldryRLM9J/jH0cVOguVXT3YjoZJsSz3F8SWKAxt2XqIxTp+eNgpBTSc+Rt3wJOSXrww1A+gL9yi6wiKYSmeuaXvzp0I9QU9fi4FMOdgSK2gmeuwzIZT5RclfBrt1QaUOdJ7/KxzrBYZ1CRDAmUvGEpKHaWJQ+cpgPwTteEPvpc71VCwH3KBJ7owh7YOCVgSjI4H8j1J7f/NoaX62NdI2DafvyWwwZ/pjm2bkcS94FrWmSYl38r/gcAw2+R7kN7swpj02kSeJvrXv3rP66BTUscdoKC1Qm+Wvr6ZArBWTEe/gK3zluaDENyCsNR0AHHn5xUECgYEA5WHPydGsQSMXfPgd5p3qxYuczmNoEdpGevfYsI5ZglHwD6KXWBDwOHIWT9zZN7rPkvis2qIeHrW/vJ086IBWWrCnJOyqWJdC79z5KsHEm3iLryaXbjCSq2jVOv9iC8lxXK+ktunnT8H7NIVJmGq/0QkgDYxEBxEi/nOJZAurzM0CgYEA7Gozvuuuojhb2hvxFvgZcKKW+rMDDylZZpdjQW7HsN3Vm1u+X3bmZgYa/OLuAqBB/wYyj5d9fE0hQFiePRanE5RGrGayb/pk93e8U3Q4jalH0HXSfOPxenOv/1GYISuU8T01uhiYyX279l1xhIgIgnaNnxXx+DMJyOJoOg+wSDECgYEAiBzveTnjNk/Fe1GcJKHmk0ySQuI3+ggFNKjnu3Ts5mkKw6xelFKQyV3hoYhSmUBQNk07/QygRObinODNKMYm9+/FbyLEvocboEwGyvlqWjtA4DFNQISwM+ikRzsOCiKCE1birevEZ86a6wKAtmH56ue4TnHCWFMx6b3TdDt8ua0CgYEAiXCkbJXvjHdfTsl3u3a4s3aSd2SinqnXQ4E3Ps78YKJWkF8hqbYh86YRAnubCwPH2k7sDfZwCVd9wT7Dn5gPwZ4uCAQcMGVNVWDSp2GljxWenl4g5GJP/HRNGQ2Cd6YyjJydI3cxO9JQ98UTPY0oBVCOvLGxXi2tvo+BMF/2h8ECgYBF6v52wt8UKV3+cTL136g+9vthlSgl/6yeLuDIhCBlmYf71K50MUYzggM06kHa54/CbLMABYcqkMsSBtcebvRNCQ4FrVIfpVlElqK4cOzNqcU0601TV4exhtUFwbGtAzkofGLrrcIiojMcNi2yGvImpF0fiuq/uggiREJ3R+Iy2w==", "QmZn1v9w5mkGMWdThF7spirifjWEa58k1zP7Ksxn4hyziG"},
	{"CAASqAkwggSkAgEAAoIBAQDULVfGUd0M2rESXYhp7e7Spzz/Hv/TRKMephSdLIbrgUjtYd4R2s5cSeIR9gcnhBCriBvi9TeBA4lDY39u682CCRtGsAfPfGOInkX5NUdF6ZQ/y2+C/GSCq/fiij8gOdPm/4twx4kzGe1xAMDOewvXxTOUWHoDx9xlz4fY2snl6LcXMhA9EabbPpyQJVq9/zikWY1ZnKnBQrzks1ALbNf60g5AYwJDxkb83r7UAIBVkpkf5OA+puYDD+YHWo7Xv+5iisfiT1szlf6WUMyfCLI/G0nENx5XTpkGSausALjl+yvwgMotEvI09LscrAl82ixuxm7402kFiHVsY+ONR7FbAgMBAAECggEAdjOkeIlWmjii+NY+jrTMnXIpmv8MCghsz0A6r60EpOJrXENvUYcR+3v1g4gKSEtcdhLnZjKx7x+nMnGVppea71xfY6vhTge/83YpCUJTKHGN81REFTbCT8G54OyfajaNBF1Ms1GV72/8RT3kK6OE33mu2G/J2Z1X/Sf2SCrq2zccFp9oMMtjURYNvIFg3BHpOztBV+O1/wFE5aO1eWAhw7Lxm0h04rtQm9h5QFP4F1z5H0OVHAW5xAoXW4zyTn/Dschxlox8XpIaO8bmdZwNSGUI5oboCXo+52ssAqY5xuzlpf224RB5KuXt7MEBS2wnI3xamUdEMoFaUlA9OsseUQKBgQDf7FJX4nTsIVBwGD94fhpt1GNCqtDlYkKLsQrfRJDGTRpXJRZwaerT/5fHsSmhUnD6qasPJzf3deUoGBILWYsggAKw1iYu/4ZuD8PhyKuZB7V9yYUa0GWy66jKlxgjqkEPfp0KYd2q9N+3l9hEVo1iizWpTn1RDxaoFhdSGuoqdQKBgQDykkXrCJ6UaR310IwNqAb4+EpYSqkPQHoD+C4I/rCvq0z1dCp6L1fpsMyWkHPyL+assa5HnxPLU27YNA15uPNQGMYUHBAiWfaRpfgjyI+XUCWhREARhRIY7YAGqoKZxNSgKXLMIQlsz0+cxwo5JkMCoXne/pc8mANYxpFXxk7SjwKBgQDdyoRw+iPXctPqg4nyfe0Nlg3PDcQlP+Mr2+sp5A8F637IjDqik7Z9zy72IztZx3+SOBXvx2e+2u5kfRr6VPHh+gkwQJzHdl8fisv5Sjr1M1aHxd/qBDqHMrYYG/pT8SHKCB3iF7doym0Auw0B+zzFO5+mF9E2RZyqVl84+uGvtQKBgAc2vGeR+Q1W1vJBgnjBf3uV/rgDeGWguEA3+D6CPITs5jwCeWTq1YV4oYz0vM3+CLEoE5PDslYApDI/0grFqk6+fd2JPB2ZaVDJi8icpVCNrKDQI8uhlnkxTvZjycCC7wPYV7akYOBghfCWpsIuVh3U6YftMFg3+RmBxj7DAgMXAoGBAN4ndnUboQqv6Seb/J4sU3NayGRvXpQlJUKi6NIjXDiDZiZQgFGA9tINHSgWVF9IEhz7510t8eJ1MEmcQFL4MnjH1XfgFnh/XBBILnSB0im2pq07UCutoywEefKb/hwI96F6/FhPzQOZu7vG1qSQoqxpkrqxv/VCxcelwo5PN56e", "QmVCuW9TryE7EMi6KDpu6fipmVp2wATp2pPunTsH86LL3k"},
	{"CAASpwkwggSjAgEAAoIBAQD2a2uSO7IbMBqy2uEjyKxehbdoxSDyEb5kvdqSQYIt5cKjw9UvRIyBY9QcH8dKSM2CVqVVgFg9Si+VkW4ABNo9w/GXVhxA5rv0MRoWGtsUUxsOGahebepSyHVUlia4atBoo2xhKBBZ4SHQIjYn+sIiDIv97tC37OycnmB1VjXMNW85etcFIgYJ6JqK6DCYL6Mojk/9V0EvdRtKzmSb32EpZtLDmcdPFAdomXTftNs3B31gKxp+MyWXIuB2GF9AVXOPvA0b70ESg62tHMFudI4Lt7jpqDgKawi1bzuo8WeQI9D0A+2gW3pVazUqw/9heANk0CutlX3K1T/Yc6lp2HoJAgMBAAECggEBANuKEKiT7rYyQVcfkn3jB06fKyx4lEWWcV4nl1e+bVCe5q4ohwI5vER6wDreRRt1iUKaF2r5kaPpOEO08Z+qsxVcQR1nZnycFNAV/vu6qXyOHgiN62dV3fIrDf1yWMTsxQi0sJZOr2KkJQ/dqworalITg5WU8vcqahOPNrjTOdLX+NalCd3XKzvzXW8R8nkLY+KkXm6K3PzBROcgCh0r36URfOEMJ2iDDke+C9uT4rWZdPXKEJGZXWICmC2qBSGraj+ZIVw1suPLGS5Ut61MO3lX7FeV3vdV+Hlnc7R4DbqbGWpVJHS3C9LPH54N6T1gbtmdljG4uvSFsficCAJeSVECgYEA/Qq3LJCXzm41JFIhzRq8RX6R0Y6xY1mF/wSnBztMTnwK/VmpotPfxsxFIjqinfoTmO9KgWshDixMohN/iYAkNBt1fJ7pxaW8A58idBrING8Vn6z6Epo4d3X6xrW4GW67LqCMK8oOIF4kv5DkXdWuQYd3cDjABGfWyL9MeOpjaZsCgYEA+Uzi2Iyz2Oza1AVkSuyG+BfplVdnkvxd+js5OvrvFympe5Lm5ec+dvc6BI8mak3Jzs1yRkefhWrfP6ETIbVi/E5cjXvHxTe4HxiP8ZhVlkA2HFNqyip/fT1rvFRZej57nCV9aOCC6N+shN8SOxfSDTu7wHUr/EQkXwI46uEvhysCgYBr93V+u1c6ikV8cNrhQ8YPNNkM/ABuLpWA4UBPUprVZhHeVbKOui5iCWh6GWCnXTudZoR4wfgBrx8njIA/cACChzDS8o82eOfG99Bgj8jarocgcVLmOw1a02kj/gdGOrv9Qh1s1bK/VswBPNZjvzex2BY7OMudVZ0MXKtJcGDofwKBgHq2SNb08zXs6I4ClWp34LvP8W9MVbQ5Ov0IF9SbXgLVxBIKrlMuL0YZzVofKHadEaAU6pIDgVcH0xob7DHkefPdkpsl+aBvG9danMf8BuztcHUi7mbS4mxARn6uOlj81pV2srNkB6wfJWlF8FBcGk3fQDyssFj9JbYKfv/GlkF/AoGAF+uC9Y533RTgyHxbCSq6NmuQ4GmMX+9Z0dwsHFXpULKct68aGVhf512DaL76EscZXrojTcx4bVmqfs2moP2fyLNiwtn90rgPTj1G5eHuJIjEsfW2H4CmJou/+QLQYVZH4830ktJK/3PnqTgNxylsr2NMUp+lUWuO63o/khDBuDg=", "QmYjNxvAYwfNz173FDhW6HSCj9jkGMnoRYtdernZLtvfrJ"},
	{"CAASqgkwggSmAgEAAoIBAQCrYpED9n2DLZ/TZEunar8uLUp6YxG22K22mCTnQsAAmRbnBApkH/rCRrrnzLz/n5yZuX+NXLh9mjzLbqVlgr9avEuETh0bWeU/gg3ts5V+qJUQARlJGj4z8qNZue42Ov4rfcyx7SieGRBAdKtR4pfaJq8rCAxMoW7JsubT2rKqw4SbBV7sDPnkAq+Bf3W4C4jfvuDjqZUeMn8ssDOMRxJ2BkZwq/TnlyxNVruG8csf9hfHgl7S3tkGNMjrNvjDWHAtXcoYZ3R+S5FpQUDQfRxjXfINv91/MOVZ0fbv5GKl+trTMo8nbyhd/8hnFNFIYxwHvJ5NR/scpbF/aZKKY3+7AgMBAAECggEBAISTEqqGGkFHIcNcaklvgcQutNZHnIRcyMPenW1nbS8JPXYm4gLh2hA+toZpEiqLRZCEte6cMvq2PSEAzmDf6zY9Qg1uf4WQ3I1sBWpC6Pm5XtWqg+zQqB13LHVUJrk8mTD7d0SLuEfUC4ZQudX1+pF+KzKWXy1IE3NjOLvZCSyeHaZJJONB/wXD7HeBE/OkQmtyBtGCYsoR1eRhvu76stUwhYSIVmipGxccr5mYsE6RXLRQqu0ZzBovszQSOCUuIq53a/Gt1eCPvAZX3FOd8IvL8+oLcYvihv/gdlZmG0qg9NFCobSxz6rzp3ZObomcv0cjIRVJ9M3UE8yKyHXww8ECgYEAyBE2OrjrsZIfQRPWRN2F8JhGcH3y3co5YUpGeAY5fik7v1bRQLNqypKJ1m26Ky5CO5UvEnnbPURpeL7QJhOUGawVOG9oceo7C9t7sMXM0ut6ObtOH+Aa00WqEm8Z+eJY/JQVy5sICnnUGfkppdxNF6QFvc7MWHFp+0/VFJSfq8UCgYEA20ySh2Tn01RWeP7IHjTZl2RVFGTb+0Ef27jYfHWQA6LxQWO2SnALj1ZpzlKpS1CWmIB3vaJP34YwPV6WrEV9R3QvGzVo0FTvX1jpdnPVEXfoPvpCarqTklRFB0A2duZP/CEuM19hzXnTw4wlj2f9P7bM2iZHjWFhwg9E4iR/tX8CgYEAvLbPilCHsfV9GauWf1rdTna0asPC5MYtncGr+ucUMBbAztkbhIY4g1/6OJND+hsmSoGL50eeL4pqWoTSXjg1iImzcopT4K/qhmK/p5zMeV+46N/u7046v3KE2+KEhBqniYg9jtJroUYNdXp7eIH3DgefisyYIAzKxU0mVsViX2kCgYEAnBKpDjcgTmTtjfs2DVlrqrU7X+JQNUcqF0Q8vock0ZG7xd/jMqL/dyn272IhHnriUvLLRWkpE1n7JxIUhdKG0L29cM3YJLzuB5vNvAHAGEiQP5H4huD2eeQDpJ9so75SGoy6xtERWB1mOvuZF7DYqem9bVxk/BfcbTJUweVcBp0CgYEAkJq00aRC6Gv7EjoHVze4VTZgAAuoWDsoLULU0a15XXrEvkoUtnDICsilDFpD1WC/GZgQcMoaRnGgdpU4oX7F2hnmzha44xGtCRWrsSjsL6tCQCJG7nxowC8O0b5N45zOCGqdm5bYx/4nH+ZdqnBWIjfjuqO1h0mO5W2FevkMp14=", "QmUsyWYq7zEj9WD4YgS653jeekAesABz4J6Tq2JsNLinan"},
	{"CAASpgkwggSiAgEAAoIBAQDR8+wPtuTQ6PqFWDp5vz097R210dXdF0yVcdvCmQ38wG339My/RTogXBMNcHYyoF7xVALHQEeGxjeEm9mhGwUV1PK4HAbuv5AKG10D92C8OVgXQt3YQi0ZcWbfas/M/bmoQw/2Gvwz6u8h4MxEYELAPzuMwuFNAF78i3g2aRnPuzkc1ozBZ9Y5z0o9m5aqZxnHuN23JHF2NWDmgVp8tsZkfXldKGq0TORBEReMJR8k3P0CygxGoa2NHaRyW86pEfRjWJA9ouTvOZ+XhghuuQ1hYds6pU7K5I9sfquKkMmmti2+mHavi7cK9taX1IbyIVp3qvBVMCk6VUVBDZW3iTCvAgMBAAECggEAGRf3WfV5KcL+1lsyOgTyc+lYSLf4wMEdJSuDoaGbe18ghadbpWzHwsBegpezeN+UGXH6FwiGxAQC9LWP8GKDXBWkoP6wkW6R77NPjb7OcBDGh5k1XLikHUthiUiB57VXsFW7naWiCS6GQF1W7ME4mTO1kWBlf0eREsw0pQQw0pQNjzHCA2CjdYL8hW8Ns9I4pgGk4M9LLlTeERdm1wByCteGJfn4Px86K/pDFJZuX03ROunJHpDn+CGqw9EQ22jWzzJFSjMZPW7X/VSHRPa+l24HGvMwpSkfiV/Wj5nI899DQq9Bn+Zc3C+DecwSLbDG4q6WGrACKiS5ydZEGzRUKQKBgQDgLzmVt9XiwybUJhHC2XmPXqD7dcISrnNF+zHLRJG4HL692AGqUdHfO6sJolWFa/PCh6003nDYx3HvbbDVGOMqg5oYGlbcXBRmGnaNcK8uT+P75NuQwjfsyJ3t0qq+0VcmkM8OWCIXXb6ykR6nU5+25UB+y+ufbjCM4LnbumXbcwKBgQDvv6bLqeNNbsy8zXYuUT2phD4gNMNndA6zBg/t61UIKMbFzHXNdF/7VyfK88uEOd51YhPaTeF2Dsr0Z+3rMz6RIu+j4k/gKYN/v+lSFzxyAIJl0sFbMPqoYD7/sIBIz2p+w9BkN+hGeAh36U+tS9q+nip9yHi0NsoHe0nCol1+1QKBgCejsgbrDoKeBwuT/6f4VCopjUVpOPucpP1GwTMz2KA5VPC0dPbsqLNUFZYKghypTdyjqNikTEfIXDj9qDrFv8UQp/qDDcDA90pme3fe6NrDfYhYmwopjUHs5x/aFB7RXRuQl0vuAXkoNkPUrRrhCmiIeLLy06LJFQST58kWWFRJAoGAfSXLyZBbCLaHxR5zTGY4C97uFx2zyrL4YVcOxaEMd//hEqtR6veOisLKENUGLXWPvDKALnps8JV0N/Rwa1AKnvRfp2qhS2AgNnVVM/bRJKlDaCeFqzZ8AZyMBih6LoenDZjllQriArbxpLZzUwC/pLUndRVKfLgE6bbfn+vlZikCgYBnXK7if2xlhfV8ESf+g/t0jHIzFjWg0bS8sPY/uMEs8PGVg4RqD9yCmgJS1GsFQSCbAhM1lVZAjIUSjgVw2T6XWTLKb86WhR7m8ViXcRZVdVpAMQXRlkcx9CWNZbXSt63VG1GgLuwXvTDxjyhUUJaX5cjiBS/dMl9dfUKUlnVv8A==", "QmQmwBU76y91Hs5xihndE7UXGQiujvNL791Yp3a58MjziX"},
	{"CAASpwkwggSjAgEAAoIBAQDACiqtbAeIR0gKZZfWuNgDssXnQnEQNrAlISlNMrtULuCtsLBk2tZ4C508T4/JQHfvbazZ/aPvkhr9KBaH8AzDU3FngHQnWblGtfm/0FAXbXPfn6DZ1rbA9rx9XpVZ+pUBDve0YxTSPOo5wOOR9u30JEvO47n1R/bF+wtMRHvDyRuoy4H86XxwMR76LYbgSlJm6SSKnrAVoWR9zqjXdaF1QljO77VbivnR5aS9vQ5Sd1mktwgb3SYUMlEGedtcMdLd3MPVCLFzq6cdjhSwVAxZ3RowR2m0hSEE/p6CKH9xz4wkMmjVrADfQTYU7spym1NBaNCrW1f+r4ScDEqI1yynAgMBAAECggEAWuJ04C5IQk654XHDMnO4h8eLsa7YI3w+UNQo38gqr+SfoJQGZzTKW3XjrC9bNTu1hzK4o1JOy4qyCy11vE/3Olm7SeiZECZ+cOCemhDUVsIOHL9HONFNHHWpLwwcUsEs05tpz400xWrezwZirSnX47tpxTgxQcwVFg2Bg07F5BntepqX+Ns7s2XTEc7YO8o77viYbpfPSjrsToahWP7ngIL4ymDjrZjgWTPZC7AzobDbhjTh5XuVKh60eUz0O7/Ezj2QK00NNkkD7nplU0tojZF10qXKCbECPn3pocVPAetTkwB1Zabq2tC2Y10dYlef0B2fkktJ4PAJyMszx4toQQKBgQD+69aoMf3Wcbw1Z1e9IcOutArrnSi9N0lVD7X2B6HHQGbHkuVyEXR10/8u4HVtbM850ZQjVnSTa4i9XJAy98FWwNS4zFh3OWVhgp/hXIetIlZF72GEi/yVFBhFMcKvXEpO/orEXMOJRdLb/7kNpMvl4MQ/fGWOmQ3InkKxLZFJ+wKBgQDA2jUTvSjjFVtOJBYVuTkfO1DKRGu7QQqNeF978ZEoU0b887kPu2yzx9pK0PzjPffpfUsa9myDSu7rncBX1FP0gNmSIAUja2pwMvJDU2VmE3Ua30Z1gVG1enCdl5ZWufum8Q+0AUqVkBdhPxw+XDJStA95FUArJzeZ2MTwbZH0RQKBgDG188og1Ys36qfPW0C6kNpEqcyAfS1I1rgLtEQiAN5GJMTOVIgF91vy11Rg2QVZrp9ryyOI/HqzAZtLraMCxWURfWn8D1RQkQCO5HaiAKM2ivRgVffvBHZd0M3NglWH/cWhxZW9MTRXtWLJX2DVvh0504s9yuAf4Jw6oG7EoAx5AoGBAJluAURO/jSMTTQB6cAmuJdsbX4+qSc1O9wJpI3LRp06hAPDM7ycdIMjwTw8wLVaG96bXCF7ZCGggCzcOKanupOP34kuCGiBkRDqt2tw8f8gA875S+k4lXU4kFgQvf8JwHi02LVxQZF0LeWkfCfw2eiKcLT4fzDV5ppzp1tREQmxAoGAGOXFomnMU9WPxJp6oaw1ZimtOXcAGHzKwiwQiw7AAWbQ+8RrL2AQHq0hD0eeacOAKsh89OXsdS9iW9GQ1mFR3FA7Kp5srjCMKNMgNSBNIb49iiG9O6P6UcO+RbYGg3CkSTG33W8l2pFIjBrtGktF5GoJudAPR4RXhVsRYZMiGag=", "QmTqawxrPeTRUKS4GSUURaC16o4etPSJv7Akq6a9xqGZUh"},
}

var peerInfo []PeerInfo

// GetTestPeerInfo gets PeerInfo for constructing a peer for testing.
func GetTestPeerInfo(i int) *PeerInfo {
	if peerInfo == nil {
		peerInfo = make([]PeerInfo, len(encoded))
	}
	info := &peerInfo[i]
	if len(info.PeerID) == 0 {
		e := encoded[i]
		data, err := base64.StdEncoding.DecodeString(e.Base64PrivKey)
		if err != nil {
			panic(err)
		}
		pk, err := crypto.UnmarshalPrivateKey(data)
		if err != nil {
			panic(err)
		}
		pid, err := peer.IDB58Decode(e.B58PeerID)
		if err != nil {
			panic(err)
		}
		info.PrivKey = pk
		info.PubKey = pk.GetPublic()
		info.PeerID = pid
		info.EncodedPrivKey = e.Base64PrivKey
		info.EncodedPeerID = e.B58PeerID
	}
	return info
}
