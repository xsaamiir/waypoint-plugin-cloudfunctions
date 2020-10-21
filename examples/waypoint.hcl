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
    use "cloudfunctions" {}
  }

  release {
    use "cloudfunctions" {}
  }
}
