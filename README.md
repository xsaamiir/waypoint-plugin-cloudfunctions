# Waypoint Plugin Google App Engine

waypoint-plugin-cloudfunctions is a deploy (platform & release) plugin for [Waypoint](https://github.com/hashicorp/waypoint).
It allows you to stage previously built zip artifcats to Google App Engine and then release the staged deployment and open it to general traffic.
The plugin is only compatible with Google App Engine Standard Environment for the time being.

**The plugin works but as expected for my use case but is still missing some features, please open an issue for any feedback, issues or missing features.**

# Install
To install the plugin, run the following command:

````bash
git clone git@github.com:sharkyze/waypoint-plugin-cloudfunctions.git # or gh repo clone sharkyze/waypoint-plugin-cloudfunctions
cd waypoint-plugin-cloudfunctions
make install
````

# Google App Engine Authentication
Please follow the instructions in the [Google Cloud Run tutorial](https://learn.hashicorp.com/tutorials/waypoint/google-cloud-run?in=waypoint/deploy-google-cloud#authenticate-to-google-cloud).
This plugin uses GCP Application Default Credentials (ADC) for authentication. More info [here](https://cloud.google.com/docs/authentication/production).

# Configure
```hcl
project = "project-name"

app "webapp" {
  path = "./webapp"
  
  url {
    auto_hostname = false
  }

  build {
    use "archive" {
      sources = ["."] # Sources are relative to /path/to/project-name/webapp/
      output_name = "function-source.zip"
      overwrite_existing = true
      ignore = [".git"]
      collapse_top_level_folder = true
    }

    registry {
      use "cloudfunctions" {}
    }

    deploy {
      use "cloudfunctions" {}
    }
    
    release {
      use "cloudfunctions" {}
    }
  }
}
```
