data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./models",
    "--dialect", "mysql", // | postgres | sqlite
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
//   dev = "mysql://root:12345@127.0.0.1:3306/hellogin3"
  dev = "docker://mysql/8/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
  diff {
    skip {
      drop_schema = false
      drop_table  = false
    }
  }
}