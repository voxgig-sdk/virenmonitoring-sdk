<?php
declare(strict_types=1);

// Virenmonitoring SDK configuration

class VirenmonitoringConfig
{
    /** @var array<string,mixed>|null */
    private static ?array $shared_config = null;

    /**
     * Return the process-wide config, built once on first use. The SDK reads
     * the config on every request and never writes to it, so one instance is
     * shared by every client rather than rebuilt per client.
     *
     * PHP arrays are copy-on-write, so callers that do mutate the result get
     * their own copy and cannot disturb the shared one.
     */
    public static function shared_config(): array
    {
        if (self::$shared_config === null) {
            self::$shared_config = self::make_config();
        }
        return self::$shared_config;
    }

    /**
     * Build a fresh, fully materialised config array. Every call rebuilds the
     * whole structure, so prefer shared_config unless you need a private copy.
     */
    public static function make_config(): array
    {
        return [
            "main" => [
                "name" => "Virenmonitoring",
                "slug" => "virenmonitoring",
                "version" => "0.0.1",
                "target" => "php",
            ],
            "feature" => [
                "test" => [
          'options' => [
            'active' => false,
          ],
          'transport' => 'base',
        ],
            ],
            "options" => [
                "base" => "https://data.bs.ch/api",
                "headers" => [
          'content-type' => 'application/json',
        ],
                "entity" => [
                    "dataset_metadata" => [],
                    "virus_monitoring" => [],
                ],
            ],
            "entity" => [
        'dataset_metadata' => [
          'fields' => [
            [
              'name' => 'description',
              'short' => 'Field description',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'label',
              'short' => 'Field label',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'short' => 'Field name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'type',
              'short' => 'Field data type',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'dataset_metadata',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'de',
                        'kind' => 'query',
                        'name' => 'lang',
                        'orig' => 'lang',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/datasets/1.0/100304/',
                  'parts' => [
                    'datasets',
                    '1.0',
                    '100304',
                  ],
                  'select' => [
                    'exist' => [
                      'lang',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'virus_monitoring' => [
          'fields' => [
            [
              'name' => 'datasetid',
              'short' => 'Dataset identifier',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'fields',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'record_timestamp',
              'short' => 'Timestamp when the record was created/updated',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'recordid',
              'short' => 'Unique record identifier',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'virus_monitoring',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => '100304',
                        'kind' => 'query',
                        'name' => 'dataset',
                        'orig' => 'dataset',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'exclude',
                        'orig' => 'exclude',
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => 'de',
                        'kind' => 'query',
                        'name' => 'lang',
                        'orig' => 'lang',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'q',
                        'orig' => 'q',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'refine',
                        'orig' => 'refine',
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => 10,
                        'kind' => 'query',
                        'name' => 'row',
                        'orig' => 'row',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'sort',
                        'orig' => 'sort',
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => 0,
                        'kind' => 'query',
                        'name' => 'start',
                        'orig' => 'start',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/records/1.0/search/',
                  'parts' => [
                    'records',
                    '1.0',
                    'search',
                  ],
                  'select' => [
                    'exist' => [
                      'dataset',
                      'exclude',
                      'lang',
                      'q',
                      'refine',
                      'row',
                      'sort',
                      'start',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
      ],
        ];
    }


    public static function make_feature(string $name)
    {
        require_once __DIR__ . '/features.php';
        return VirenmonitoringFeatures::make_feature($name);
    }
}
