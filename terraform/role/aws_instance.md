resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ec2:AuthorizeSecurityGroupEgress"
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
  tags = { createdby = "JamesWoolfenden" }
}

resource "aws_iam_role_policy_attachment" "basic" {
  role       = aws_iam_role.basic.name
  policy_arn = aws_iam_policy.basic.arn
}

resource "aws_iam_user_policy_attachment" "basic" {
  user       = "basic"
  policy_arn = aws_iam_policy.basic.arn
}