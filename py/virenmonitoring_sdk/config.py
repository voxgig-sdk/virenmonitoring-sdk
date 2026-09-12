# Virenmonitoring SDK configuration


# The sekreto plugin DEFINITIONS the model selected per feature, imported
# above by name from the modules the catalogue's active `plugin.def`
# entries declare. Handed to each feature (secrets builds its Sekreto
# with them): a provider kind not listed here is unknown to that SDK.
FEATURE_PLUGINS = {
}


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "Virenmonitoring",
            "slug": "virenmonitoring",
            "version": "0.0.1",
            "target": "py",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
        "transport": "base",
      },
        },
        "options": {
            "base": "https://data.bs.ch/api",
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "dataset_metadata": {},
                "virus_monitoring": {},
            },
        },
        "entity": {
      "dataset_metadata": {
        "fields": [
          {
            "name": "description",
            "short": "Field description",
            "type": "`$STRING`",
          },
          {
            "name": "label",
            "short": "Field label",
            "type": "`$STRING`",
          },
          {
            "name": "name",
            "short": "Field name",
            "type": "`$STRING`",
          },
          {
            "name": "type",
            "short": "Field data type",
            "type": "`$STRING`",
          },
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
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/datasets/1.0/100304/",
                "segments": [
                  {
                    "lit": "datasets",
                  },
                  {
                    "lit": "1.0",
                  },
                  {
                    "lit": "100304",
                  },
                ],
                "select": {
                  "exist": [
                    "lang",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "datasets",
                  "1.0",
                  "100304",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "virus_monitoring": {
        "fields": [
          {
            "name": "datasetid",
            "short": "Dataset identifier",
            "type": "`$STRING`",
          },
          {
            "name": "fields",
            "type": "`$OBJECT`",
          },
          {
            "format": "date-time",
            "name": "record_timestamp",
            "short": "Timestamp when the record was created/updated",
            "type": "`$STRING`",
          },
          {
            "name": "recordid",
            "short": "Unique record identifier",
            "type": "`$STRING`",
          },
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
                      "type": "`$STRING`",
                    },
                    {
                      "kind": "query",
                      "name": "exclude",
                      "orig": "exclude",
                      "type": "`$STRING`",
                    },
                    {
                      "example": "de",
                      "kind": "query",
                      "name": "lang",
                      "orig": "lang",
                      "type": "`$STRING`",
                    },
                    {
                      "kind": "query",
                      "name": "q",
                      "orig": "q",
                      "type": "`$STRING`",
                    },
                    {
                      "kind": "query",
                      "name": "refine",
                      "orig": "refine",
                      "type": "`$STRING`",
                    },
                    {
                      "example": 10,
                      "kind": "query",
                      "name": "row",
                      "orig": "row",
                      "type": "`$INTEGER`",
                    },
                    {
                      "kind": "query",
                      "name": "sort",
                      "orig": "sort",
                      "type": "`$STRING`",
                    },
                    {
                      "example": 0,
                      "kind": "query",
                      "name": "start",
                      "orig": "start",
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/records/1.0/search/",
                "segments": [
                  {
                    "lit": "records",
                  },
                  {
                    "lit": "1.0",
                  },
                  {
                    "lit": "search",
                  },
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
                    "start",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "records",
                  "1.0",
                  "search",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
