project = "examples"

app "hello-http" {
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

app "hello-pubsub" {
  path = "./helloworld"

  url {
    auto_hostname = false
  }

  build {
    use "archive" {}

    registry {
      use "cloudfunctions" {
        project = "project-name"
        location = "europe-west1"
      }
    }
  }

  deploy {
    use "cloudfunctions" {
      entry_point = "HelloPubSub"
      description = "Deployed using Waypoint 🎉"
      runtime = "go113"
      max_instances = 1
      available_memory_mb = 128
      event_trigger {
        event_type = "google.pubsub.topic.publish"
        resource = "projects/project-name/topics/topic-name"
        failure_policy {
          retry = true
        }
      }
    }
  }

  release {
    use "cloudfunctions" {}
  }
}
