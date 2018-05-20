resource "aws_ecs_service" "golang-fibonacci" {
  name            = "golang-fibonacci"
  cluster         = "test"
  task_definition = "${aws_ecs_task_definition.golang-fibonacci.arn}"
  desired_count   = 1
  iam_role        = ""

  ordered_placement_strategy {
    type  = "binpack"
    field = "cpu"
  }

}

