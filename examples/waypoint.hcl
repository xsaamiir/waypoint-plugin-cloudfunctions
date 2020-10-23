project="examples"

app "helloworld" {
  path = "./helloworld"

  build {
    use "archive" {}

    registry {
      use "cloudfunctions" {
        project = "project-id"
        location = "europe-west1"
      }
    }
  }

  deploy {
    use "cloudfunctions" {
      entry_point = "HelloHTTP"
      description = "Deployed using Waypoint 🎉"
      runtime = "go113"
      max_instances = 1
      available_memory_mb = 128
      trigger_http = true
    }
  }

  release {
    use "cloudfunctions" {
      unauthenticated = true
    }
  }
}
