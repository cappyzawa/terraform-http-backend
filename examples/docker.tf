terraform {
  backend "http" {
    address        = "http://localhost:8080/"
    lock_address   = "http://localhost:8080/lock"
    lock_method    = "POST"
    unlock_address = "http://localhost:8080/lock"
    unlock_method  = "DELETE"
  }
}

resource "docker_image" "nginx" {
  name = "nginx:latest"
}

resource "docker_container" "nginx" {
  image = docker_image.nginx.latest
  name  = "tutorial"
  ports {
    internal = 80
    external = 80
  }
}
