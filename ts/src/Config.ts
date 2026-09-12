
import { BaseFeature } from './feature/base/BaseFeature'
import { TestFeature } from './feature/test/TestFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   test: TestFeature,

}


// Per-feature plugin DEFINITIONS (voxgig/plugin `Definition` values), from
// the model's active plugin groups. A feature that takes a `plugins` option
// (secrets over sekreto) reads its own entry; a feature with no plugins has
// none. Named imports above make each definition statically reachable, so
// an SDK carries exactly the plugin modules its model selects — the same
// leanness the old side-effect registry imports bought, without a registry.
const FEATURE_PLUGINS: Record<string, any[]> = {
  
}


class Config {

  makeFeature(this: any, fn: string) {
    const fc = FEATURE_CLASS[fn]
    const fi = new fc()
    // TODO: errors etc
    return fi
  }

  // False for a feature added at runtime via options.extend (station's
  // adopt path) - the constructor uses this to skip makeFeature for names
  // no generated class backs.
  hasFeature(this: any, fn: string) {
    return null != FEATURE_CLASS[fn]
  }


  main = {
    name: 'Virenmonitoring',
        slug: "virenmonitoring",
    version: "0.0.1",
    target: "ts",

  }


  feature = {
     test:     {
      "options": {
        "active": false
      },
      "transport": "base"
    },

  }


  options = {
    base: "https://data.bs.ch/api",

    headers: {
      "content-type": "application/json"
    },

    entity: {
      
      dataset_metadata: {
      },

      virus_monitoring: {
      },

    }
  }


  entity = {
    "dataset_metadata": {
      "fields": [
        {
          "name": "description",
          "short": "Field description",
          "type": "`$STRING`"
        },
        {
          "name": "label",
          "short": "Field label",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "short": "Field name",
          "type": "`$STRING`"
        },
        {
          "name": "type",
          "short": "Field data type",
          "type": "`$STRING`"
        }
      ],
      "name": "dataset_metadata",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "example": "de",
                    "kind": "query",
                    "name": "lang",
                    "orig": "lang",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/datasets/1.0/100304/",
              "segments": [
                {
                  "lit": "datasets"
                },
                {
                  "lit": "1.0"
                },
                {
                  "lit": "100304"
                }
              ],
              "select": {
                "exist": [
                  "lang"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "parts": [
                "datasets",
                "1.0",
                "100304"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "virus_monitoring": {
      "fields": [
        {
          "name": "datasetid",
          "short": "Dataset identifier",
          "type": "`$STRING`"
        },
        {
          "name": "fields",
          "type": "`$OBJECT`"
        },
        {
          "format": "date-time",
          "name": "record_timestamp",
          "short": "Timestamp when the record was created/updated",
          "type": "`$STRING`"
        },
        {
          "name": "recordid",
          "short": "Unique record identifier",
          "type": "`$STRING`"
        }
      ],
      "name": "virus_monitoring",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "example": "100304",
                    "kind": "query",
                    "name": "dataset",
                    "orig": "dataset",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "exclude",
                    "orig": "exclude",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "de",
                    "kind": "query",
                    "name": "lang",
                    "orig": "lang",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "q",
                    "orig": "q",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "refine",
                    "orig": "refine",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 10,
                    "kind": "query",
                    "name": "row",
                    "orig": "row",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 0,
                    "kind": "query",
                    "name": "start",
                    "orig": "start",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/records/1.0/search/",
              "segments": [
                {
                  "lit": "records"
                },
                {
                  "lit": "1.0"
                },
                {
                  "lit": "search"
                }
              ],
              "select": {
                "exist": [
                  "dataset",
                  "exclude",
                  "lang",
                  "q",
                  "refine",
                  "row",
                  "sort",
                  "start"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "parts": [
                "records",
                "1.0",
                "search"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    }
  }
}


const config = new Config()

export {
  config,
  FEATURE_PLUGINS,
}

