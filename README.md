# Waypoint Plugin Google Cloud Functions

waypoint-plugin-cloudfunctions is a deploy (registry, platform & release) plugin
for [Waypoint](https://github.com/hashicorp/waypoint). It allows you to stage previously built zip artifcats to Google
Cloud Functions and then release the staged deployment and open it to general traffic.

# Status

I am currently using this plugin to deploy several Google Cloud Functions in production. It is working fine as far as my
use case goes. Please open an issue for feedback, issues or feature requests.

## Current Limitation

- The plugin doesn't support staging deployments before releasing them to general traffic, this is mainly because I
  haven't found a way to support this using Cloud Functions. This means that the plugin does almost nothing in
  the `release` stage expect setting the IAM policy for unauthenticated functions.
- `Destroy` is not supported for `build` nor `release` for the same reason mentioned above. The only way I was able to
  implement this was by deleting the function, which is not what most people would want I think.

# Install

To install the plugin, run the following command:

````bash
# First install the plugin archive by following the instuction here: https://github.com/sharkyze/waypoint-plugin-archive
git clone git@github.com:sharkyze/waypoint-plugin-cloudfunctions.git # or gh repo clone sharkyze/waypoint-plugin-cloudfunctions
cd waypoint-plugin-cloudfunctions
make install
````

# Authentication

Please follow the instructions in
the [Google Cloud Run tutorial](https://learn.hashicorp.com/tutorials/waypoint/google-cloud-run?in=waypoint/deploy-google-cloud#authenticate-to-google-cloud)
. This plugin uses GCP Application Default Credentials (ADC) for authentication. More
info [here](https://cloud.google.com/docs/authentication/production).

# Documentation

The documentation of the plugin is [here](./doc/README.md)

# Example

```hcl
project = "examples"

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
      ingress_settings = "ALLOW_ALL"
      trigger_http = true
    }
  }

  release {
    use "cloudfunctions" {
      unauthenticated = true
    }
  }
}
```
