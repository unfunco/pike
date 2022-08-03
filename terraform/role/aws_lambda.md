resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
        "ec2:DescribeAccountAttributes",
        "lambda:CreateFunction",
        "iam:PassRole",
        "lambda:GetFunction",
        "lambda:ListVersionsByFunction",
        "lambda:DeleteFunction",
        "lambda:GetFunctionCodeSigningConfig",
        "lambda:TagResource",
        "lambda:UntagResource"
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
	# checkov:skip=CKV_AWS_40: By design
  user       = "basic"
  policy_arn = aws_iam_policy.basic.arn
}