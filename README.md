# Status Page

![Build Status](https://github.com/bhanurp/status-page/actions/workflows/go.yml/badge.svg)
![Test Status](https://github.com/bhanurp/status-page/actions/workflows/tests.yml/badge.svg)

This repository contains the code for the status page http go client, which allows for the creation, updating, and deletion of incidents.

## Overview

The `status-page` provides an API for managing incidents on a status page. It includes functionality for
  - creating incidents
  - updating incidents
  - deleting incidents
  - fetching all incidents

## Setup

### Prerequisites

- Go 1.23 or higher
- Environment variables:
  - `STATUS_PAGE_BEARER_TOKEN`: Your API key for authentication
  - `STATUS_PAGE_ID`: The ID of your status page
  - `STATUS_PAGE_COMPONENT_ID`: The ID of your status page component

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/bhanurp/status-page.git
    cd status-page
    ```

2. Install Dependencies

    ```sh
    go mod tidy
    ```

3. Running Tests

  To run the tests, ensure that the necessary environment variables are set and use the following command:
    ```sh
    go test -v ./...
    ```

### Usage
#### Creating an Incident

To create an incident, use the CreateIncident function from the incident package. Here is an example:

```go
package main

import (
    "log"
    "os"

    "github.com/bhanurp/status-page/incident"
)

func main() {
    apiKey := os.Getenv("STATUS_PAGE_BEARER_TOKEN")
    statusPageID := os.Getenv("STATUS_PAGE_ID")
    statusPageComponentID := os.Getenv("STATUS_PAGE_COMPONENT_ID")

    client := incident.NewClient(apiKey, statusPageComponentID, statusPageID)
    incident := incident.Incident{
        Name:   "Test Incident",
        Status: "investigating",
    }

    createdIncident, err := client.CreateIncident(incident)
    if err != nil {
        log.Fatalf("Failed to create incident: %v", err)
    }

    log.Printf("Created incident: %v", createdIncident)
}
```

#### Updating an Incident

To update an incident, use the UpdateIncident function from the incident package. Here is an example:

```go
package main

import (
    "log"
    "os"

    "github.com/bhanurp/status-page/incident"
)

func main() {
    apiKey := os.Getenv("STATUS_PAGE_BEARER_TOKEN")
    statusPageID := os.Getenv("STATUS_PAGE_ID")
    statusPageComponentID := os.Getenv("STATUS_PAGE_COMPONENT_ID")

    client := incident.NewClient(apiKey, statusPageComponentID, statusPageID)
    incidentID := "incident_id_here"
    updatedIncident := incident.Incident{
        ID:     incidentID,
        Name:   "Updated Test Incident",
        Status: "resolved",
    }

    err := client.UpdateIncident(incidentID, updatedIncident)
    if err != nil {
        log.Fatalf("Failed to update incident: %v", err)
    }

    log.Printf("Updated incident: %v", updatedIncident)
}
```

#### Deleting an Incident

To delete an incident, use the DeleteIncident function from the incident package. Here is an example:

```go
package main

import (
    "log"
    "os"

    "github.com/bhanurp/status-page/incident"
)

func main() {
    apiKey := os.Getenv("STATUS_PAGE_BEARER_TOKEN")
    statusPageID := os.Getenv("STATUS_PAGE_ID")
    statusPageComponentID := os.Getenv("STATUS_PAGE_COMPONENT_ID")

    client := incident.NewClient(apiKey, statusPageComponentID, statusPageID)
    incidentID := "incident_id_here"

    err := client.DeleteIncident(incidentID)
    if err != nil {
        log.Fatalf("Failed to delete incident: %v", err)
    }

    log.Println("Deleted incident successfully")
}
```

### Contributing
We welcome contributions! Please see CONTRIBUTING.md for details on how to contribute to this project.

### License
This project is licensed under the MIT License - see the LICENSE file for details.