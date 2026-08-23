-- Virenmonitoring SDK configuration

-- Build a fresh, fully materialised config table. Every call rebuilds the
-- whole structure, so prefer require("config_shared") unless you need a
-- private copy you intend to mutate.
local function make_config()
  return {
    main = {
      name = "Virenmonitoring",
      slug = "virenmonitoring",
      version = "0.0.1",
      target = "lua",
    },
    feature = {
      ["test"] = {
        ["options"] = {
          ["active"] = false,
        },
      },
    },
    options = {
      base = "https://data.bs.ch/api",
      headers = {
        ["content-type"] = "application/json",
      },
      entity = {
        ["dataset_metadata"] = {},
        ["virus_monitoring"] = {},
      },
    },
    entity = {
      ["dataset_metadata"] = {
        ["fields"] = {
          {
            ["name"] = "description",
            ["short"] = "Field description",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "label",
            ["short"] = "Field label",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "name",
            ["short"] = "Field name",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "type",
            ["short"] = "Field data type",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "dataset_metadata",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["example"] = "de",
                      ["kind"] = "query",
                      ["name"] = "lang",
                      ["orig"] = "lang",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/datasets/1.0/100304/",
                ["parts"] = {
                  "datasets",
                  "1.0",
                  "100304",
                },
                ["select"] = {
                  ["exist"] = {
                    "lang",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["virus_monitoring"] = {
        ["fields"] = {
          {
            ["name"] = "datasetid",
            ["short"] = "Dataset identifier",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fields",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "record_timestamp",
            ["short"] = "Timestamp when the record was created/updated",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "recordid",
            ["short"] = "Unique record identifier",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "virus_monitoring",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["example"] = "100304",
                      ["kind"] = "query",
                      ["name"] = "dataset",
                      ["orig"] = "dataset",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "exclude",
                      ["orig"] = "exclude",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = "de",
                      ["kind"] = "query",
                      ["name"] = "lang",
                      ["orig"] = "lang",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "q",
                      ["orig"] = "q",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "refine",
                      ["orig"] = "refine",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = 10,
                      ["kind"] = "query",
                      ["name"] = "row",
                      ["orig"] = "row",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = 0,
                      ["kind"] = "query",
                      ["name"] = "start",
                      ["orig"] = "start",
                      ["type"] = "`$INTEGER`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/records/1.0/search/",
                ["parts"] = {
                  "records",
                  "1.0",
                  "search",
                },
                ["select"] = {
                  ["exist"] = {
                    "dataset",
                    "exclude",
                    "lang",
                    "q",
                    "refine",
                    "row",
                    "sort",
                    "start",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
    },
  }
end


local function make_feature(name)
  local features = require("features")
  local factory = features[name]
  if factory ~= nil then
    return factory()
  end
  return features.base()
end


-- Attach make_feature to the SDK class
local function setup_sdk(SDK)
  SDK._make_feature = make_feature
end


return make_config
