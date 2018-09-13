
variable "image_name" {
    default = "josephmuli/simple_nginx:0.0.2"
}
variable "container_name" {
    default = "kilimanjaro_container"
}
variable "internal_port" {
    default = "8080"
}
variable "external_port" {
    default = "8080"
}
