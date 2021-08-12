package amorph

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/itchyny/gojq"
)

var ErrAmorphQueryError = errors.New("Amorph query error")

func Unmarshal(in Amorph, dest interface{}) (err error) {
	rdr, wrt := io.Pipe()

	decoder := json.NewDecoder(rdr)

	res := make(chan error)
	go func(w io.WriteCloser) {
		encoder := json.NewEncoder(w)
		encoder.Encode(in)
		res <- w.Close()
	}(wrt)
	e0 := decoder.Decode(dest)
	e1 := <-res
	if e0 != nil {
		err = e0
	}
	if e1 != nil {
		err = e1
	}
	return //
}

func Query(in Amorph, q string) (out Amorph, err error) {
	var val interface{}

	query, err := gojq.Parse(q)
	if err != nil {
		return //
	}
	iter := query.Run(in)
	val, ok := iter.Next()
	if !ok {
		return nil, ErrAmorphQueryError
	}
	switch tval := val.(type) {
	case nil:
		return nil, ErrAmorphQueryError
	case error:
		err = tval
		return //
	case map[string]interface{}:
		out = tval
	case []interface{}:
		out = tval
	case float64:
		out = tval // TODO we don't want these
	case string:
		out = tval
	default:
		panic("unexpected type")
	}
	if !ok {
		return nil, ErrAmorphQueryError
	}
	return //
}

func QueryAndUnmarshall(in Amorph, q string, dest interface{}) (err error) {
	sub, err := Query(in, q)
	if err != nil {
		return //
	}
	return Unmarshal(sub, dest)
}

/*
{
  "platform": {
	  "type": "drop",
	  "slug": "plat00",
          "domain": "googular.com",
          "subdomain": "clu1"
  },
  "clusters": [
    {
      "name": "TestCluster00",
      "slug": "clust00"
    }
  ],
  "instances": [
    {
      "name": "TestInstance00",
      "slug": "inst00"
    }
  ],
  "nodes": [
    {
      "name": "node0.slllc.net",
      "slug": "node0",
      "addr": "",
      "uuid": "dbe23d8e-8e9f-11eb-8dcd-0242ac130003",
      "cluster": "clust00",
      "roles": [
        {
          "role": "LEADERMON",
          "slug": "0-3",
          "config": {
            "ACRUD": [  "http://199.217.116.134:2379" , "http://69.64.50.140:2379" , "http://199.217.117.29:2379"]
,
	    "serverUrls": [
		"199.217.116.134:2920"
	    ],
	    "ACRUDDir": "team000",
            "num": "0",
            "of": "3"
          }
        }
        , {
          "role": "ROSTERMON",
          "slug": "rmon",
	  "config": {
             "ACRUD": [  "http://199.217.116.134:2379" , "http://69.64.50.140:2379" , "http://199.217.117.29:2379"]
,
             "leaders": [  "http://199.217.116.134:2379" , "http://69.64.50.140:2379" , "http://199.217.117.29:2379"]
,
	     "ACRUDDir": "team000",
	     "server": {
	        "Url": "199.217.116.134:5960"
	     },
	     "database": {
		     "store": "/var/lib/rostermon/rostermon.db"
	     }
	  }
        }
        , {
          "roleslug": "cdvros-0",
          "slug": "node0-0",
          "config": {
	    "UUID": "61040bf4-9270-11eb-a8b3-0242ac130003",
            "ACRUD": [  "http://199.217.116.134:2379" , "http://69.64.50.140:2379" , "http://199.217.117.29:2379"]
,
	    "ACRUDDir": "team000",
            "PathPrefix": "/clu1",
            "HTTPDataServer": [
              {
                "Hostname": "199.217.116.134",
                "Hostport": "8080",
                "type": "public"
              }
            ],
            "HTTPLostAndFoundServer": {
              "Hostname": "199.217.116.134",
              "Hostport": "911"
            },
            "PathMangler": {
              "URIPrefix": "/clu1",
              "MetadataRoot": "/data/clu1/meta",
              "DataRoot": "/data/clu1/data",
              "FoldedRoot": "/data/clu1/folded",
              "UnfoldedRoot": "/data/clu1/unfolded",
              "UnfoldedSuffix": "CopyIdx/%d",
              "UnfoldedReString": "(.*)/CopyIdx/(\\\\d\\\\d*)/(.*)"
            },
            "ObjectConfig": {
              "WatchdogInterval": "123456s",
              "OrphanedInterval": "720s"
            },
            "CDVRConfig": {
              "ArchiveInterval": "15s",
              "IngestWorkerHeartBeat": "600s",
              "IngestWorkers": 3,
              "MaxBatchSize": 250,
	      "FailureDomainType": "tag",
              "FailureDomainLocation": ["stl", "node0", "disk0", "part0"],
              "FailureDomainLayers": ["region", "chassis", "device", "partition"]
            },
            "BEREOSConfig": {
              "Ops": {
                "Workers": 3
              }
            }
          }
         }
          , {
          "roleslug": "cdvros-1",
          "slug": "node0-1",
          "config": {
	    "UUID": "61040cf8-9270-11eb-a8b3-0242ac130003",
            "ACRUD": [  "http://199.217.116.134:2379" , "http://69.64.50.140:2379" , "http://199.217.117.29:2379"]
,
	    "ACRUDDir": "team000",
            "PathPrefix": "/clu1",
            "HTTPDataServer": [
              {
                "Hostname": "199.217.116.134",
                "Hostport": "8081",
                "type": "public"
              }
            ],
            "HTTPLostAndFoundServer": {
              "Hostname": "199.217.116.134",
              "Hostport": "911"
            },
            "PathMangler": {
              "URIPrefix": "/clu1",
              "MetadataRoot": "/mnt/disk1/clu1/meta",
              "DataRoot": "/mnt/disk1/clu1/data",
              "FoldedRoot": "/mnt/disk1/clu1/folded",
              "UnfoldedRoot": "/mnt/disk1/clu1/unfolded",
              "UnfoldedSuffix": "CopyIdx/%d",
              "UnfoldedReString": "(.*)/CopyIdx/(\\\\d\\\\d*)/(.*)"
            },
            "ObjectConfig": {
              "WatchdogInterval": "123456s",
              "OrphanedInterval": "720s"
            },
            "CDVRConfig": {
              "ArchiveInterval": "15s",
              "IngestWorkerHeartBeat": "600s",
              "IngestWorkers": 3,
              "MaxBatchSize": 250,
	      "FailureDomainType": "tag",
              "FailureDomainLocation": ["stl", "node0", "disk1", "part0"],
              "FailureDomainLayers": ["region", "chassis", "device", "partition"]
            },
            "BEREOSConfig": {
              "Ops": {
                "Workers": 3
              }
            }
          }
         }
      ],
      "LogWriter": {
        "Hostname" : "node0.slllc.net`",
        "Nodeslug" : "node0",
        "WrtMech" : "STDOUT",
        "Money" : "dbe23d8e-8e9f-11eb-8dcd-0242ac130003"
      }
    } ,     {
      "name": "node1.slllc.net",
      "slug": "node1",
      "addr": "",
      "uuid": "dbe23d8e-8e9f-11eb-8dcd-0242ac130003",
      "cluster": "clust00",
      "roles": [
        {
          "role": "LEADERMON",
          "slug": "1-3",
          "config": {
            "ACRUD": [  "http://199.217.116.134:2379" , "http://69.64.50.140:2379" , "http://199.217.117.29:2379"]
,
	    "serverUrls": [
		"69.64.50.140:2920"
	    ],
	    "ACRUDDir": "team000",
            "num": "1",
            "of": "3"
          }
        }
        , {
          "role": "ROSTERMON",
          "slug": "rmon",
	  "config": {
             "ACRUD": [  "http://199.217.116.134:2379" , "http://69.64.50.140:2379" , "http://199.217.117.29:2379"]
,
             "leaders": [  "http://199.217.116.134:2379" , "http://69.64.50.140:2379" , "http://199.217.117.29:2379"]
,
	     "ACRUDDir": "team000",
	     "server": {
	        "Url": "69.64.50.140:5960"
	     },
	     "database": {
		     "store": "/var/lib/rostermon/rostermon.db"
	     }
	  }
        }
        , {
          "roleslug": "cdvros-0",
          "slug": "node1-0",
          "config": {
	    "UUID": "610404ce-9270-11eb-a8b3-0242ac130003",
            "ACRUD": [  "http://199.217.116.134:2379" , "http://69.64.50.140:2379" , "http://199.217.117.29:2379"]
,
	    "ACRUDDir": "team000",
            "PathPrefix": "/clu1",
            "HTTPDataServer": [
              {
                "Hostname": "69.64.50.140",
                "Hostport": "8080",
                "type": "public"
              }
            ],
            "HTTPLostAndFoundServer": {
              "Hostname": "69.64.50.140",
              "Hostport": "911"
            },
            "PathMangler": {
              "URIPrefix": "/clu1",
              "MetadataRoot": "/data/clu1/meta",
              "DataRoot": "/data/clu1/data",
              "FoldedRoot": "/data/clu1/folded",
              "UnfoldedRoot": "/data/clu1/unfolded",
              "UnfoldedSuffix": "CopyIdx/%d",
              "UnfoldedReString": "(.*)/CopyIdx/(\\\\d\\\\d*)/(.*)"
            },
            "ObjectConfig": {
              "WatchdogInterval": "123456s",
              "OrphanedInterval": "720s"
            },
            "CDVRConfig": {
              "ArchiveInterval": "15s",
              "IngestWorkerHeartBeat": "600s",
              "IngestWorkers": 3,
              "MaxBatchSize": 250,
	      "FailureDomainType": "tag",
              "FailureDomainLocation": ["stl", "node0", "disk0", "part0"],
              "FailureDomainLayers": ["region", "chassis", "device", "partition"]
            },
            "BEREOSConfig": {
              "Ops": {
                "Workers": 3
              }
            }
          }
         }
          , {
          "roleslug": "cdvros-1",
          "slug": "node1-1",
          "config": {
	    "UUID": "610406e0-9270-11eb-a8b3-0242ac130003",
            "ACRUD": [  "http://199.217.116.134:2379" , "http://69.64.50.140:2379" , "http://199.217.117.29:2379"]
,
	    "ACRUDDir": "team000",
            "PathPrefix": "/clu1",
            "HTTPDataServer": [
              {
                "Hostname": "69.64.50.140",
                "Hostport": "8081",
                "type": "public"
              }
            ],
            "HTTPLostAndFoundServer": {
              "Hostname": "69.64.50.140",
              "Hostport": "911"
            },
            "PathMangler": {
              "URIPrefix": "/clu1",
              "MetadataRoot": "/mnt/disk1/clu1/meta",
              "DataRoot": "/mnt/disk1/clu1/data",
              "FoldedRoot": "/mnt/disk1/clu1/folded",
              "UnfoldedRoot": "/mnt/disk1/clu1/unfolded",
              "UnfoldedSuffix": "CopyIdx/%d",
              "UnfoldedReString": "(.*)/CopyIdx/(\\\\d\\\\d*)/(.*)"
            },
            "ObjectConfig": {
              "WatchdogInterval": "123456s",
              "OrphanedInterval": "720s"
            },
            "CDVRConfig": {
              "ArchiveInterval": "15s",
              "IngestWorkerHeartBeat": "600s",
              "IngestWorkers": 3,
              "MaxBatchSize": 250,
	      "FailureDomainType": "tag",
              "FailureDomainLocation": ["stl", "node0", "disk1", "part0"],
              "FailureDomainLayers": ["region", "chassis", "device", "partition"]
            },
            "BEREOSConfig": {
              "Ops": {
                "Workers": 3
              }
            }
          }
         }
      ],
      "LogWriter": {
        "Hostname" : "node1.slllc.net`",
        "Nodeslug" : "node1",
        "WrtMech" : "STDOUT",
        "Money" : "dbe23d8e-8e9f-11eb-8dcd-0242ac130003"
      }
    } ,     {
      "name": "node2.slllc.net",
      "slug": "node2",
      "addr": "",
      "uuid": "dbe23d8e-8e9f-11eb-8dcd-0242ac130003",
      "cluster": "clust00",
      "roles": [
        {
          "role": "LEADERMON",
          "slug": "2-3",
          "config": {
            "ACRUD": [  "http://199.217.116.134:2379" , "http://69.64.50.140:2379" , "http://199.217.117.29:2379"]
,
	    "serverUrls": [
		"199.217.117.29:2920"
	    ],
	    "ACRUDDir": "team000",
            "num": "2",
            "of": "3"
          }
        }
        , {
          "role": "ROSTERMON",
          "slug": "rmon",
	  "config": {
             "ACRUD": [  "http://199.217.116.134:2379" , "http://69.64.50.140:2379" , "http://199.217.117.29:2379"]
,
             "leaders": [  "http://199.217.116.134:2379" , "http://69.64.50.140:2379" , "http://199.217.117.29:2379"]
,
	     "ACRUDDir": "team000",
	     "server": {
	        "Url": "199.217.117.29:5960"
	     },
	     "database": {
		     "store": "/var/lib/rostermon/rostermon.db"
	     }
	  }
        }
        , {
          "roleslug": "cdvros-0",
          "slug": "node2-0",
          "config": {
	    "UUID": "61040e2e-9270-11eb-a8b3-0242ac130003",
            "ACRUD": [  "http://199.217.116.134:2379" , "http://69.64.50.140:2379" , "http://199.217.117.29:2379"]
,
	    "ACRUDDir": "team000",
            "PathPrefix": "/clu1",
            "HTTPDataServer": [
              {
                "Hostname": "199.217.117.29",
                "Hostport": "8080",
                "type": "public"
              }
            ],
            "HTTPLostAndFoundServer": {
              "Hostname": "199.217.117.29",
              "Hostport": "911"
            },
            "PathMangler": {
              "URIPrefix": "/clu1",
              "MetadataRoot": "/data/clu1/meta",
              "DataRoot": "/data/clu1/data",
              "FoldedRoot": "/data/clu1/folded",
              "UnfoldedRoot": "/data/clu1/unfolded",
              "UnfoldedSuffix": "CopyIdx/%d",
              "UnfoldedReString": "(.*)/CopyIdx/(\\\\d\\\\d*)/(.*)"
            },
            "ObjectConfig": {
              "WatchdogInterval": "123456s",
              "OrphanedInterval": "720s"
            },
            "CDVRConfig": {
              "ArchiveInterval": "15s",
              "IngestWorkerHeartBeat": "600s",
              "IngestWorkers": 3,
              "MaxBatchSize": 250,
	      "FailureDomainType": "tag",
              "FailureDomainLocation": ["stl", "node0", "disk0", "part0"],
              "FailureDomainLayers": ["region", "chassis", "device", "partition"]
            },
            "BEREOSConfig": {
              "Ops": {
                "Workers": 3
              }
            }
          }
         }
          , {
          "roleslug": "cdvros-1",
          "slug": "node2-1",
          "config": {
	    "UUID": "61040eec-9270-11eb-a8b3-0242ac130003",
            "ACRUD": [  "http://199.217.116.134:2379" , "http://69.64.50.140:2379" , "http://199.217.117.29:2379"]
,
	    "ACRUDDir": "team000",
            "PathPrefix": "/clu1",
            "HTTPDataServer": [
              {
                "Hostname": "199.217.117.29",
                "Hostport": "8081",
                "type": "public"
              }
            ],
            "HTTPLostAndFoundServer": {
              "Hostname": "199.217.117.29",
              "Hostport": "911"
            },
            "PathMangler": {
              "URIPrefix": "/clu1",
              "MetadataRoot": "/mnt/disk1/clu1/meta",
              "DataRoot": "/mnt/disk1/clu1/data",
              "FoldedRoot": "/mnt/disk1/clu1/folded",
              "UnfoldedRoot": "/mnt/disk1/clu1/unfolded",
              "UnfoldedSuffix": "CopyIdx/%d",
              "UnfoldedReString": "(.*)/CopyIdx/(\\\\d\\\\d*)/(.*)"
            },
            "ObjectConfig": {
              "WatchdogInterval": "123456s",
              "OrphanedInterval": "720s"
            },
            "CDVRConfig": {
              "ArchiveInterval": "15s",
              "IngestWorkerHeartBeat": "600s",
              "IngestWorkers": 3,
              "MaxBatchSize": 250,
	      "FailureDomainType": "tag",
              "FailureDomainLocation": ["stl", "node0", "disk1", "part0"],
              "FailureDomainLayers": ["region", "chassis", "device", "partition"]
            },
            "BEREOSConfig": {
              "Ops": {
                "Workers": 3
              }
            }
          }
         }
      ],
      "LogWriter": {
        "Hostname" : "node2.slllc.net`",
        "Nodeslug" : "node2",
        "WrtMech" : "STDOUT",
        "Money" : "dbe23d8e-8e9f-11eb-8dcd-0242ac130003"
      }
    }   ]
}

*/
