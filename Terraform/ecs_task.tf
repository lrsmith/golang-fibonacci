resource "aws_ecs_task_definition" "golang-fibonacci" {
  family                = "golang-fibonacci"
  container_definitions = "${file("task-definitions/golang-fibonacci.json")}"

}

