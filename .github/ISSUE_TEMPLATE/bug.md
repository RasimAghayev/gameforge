name: Bug report
description: File a bug report
title: "[bug] "
body:
  - type: textarea
    attributes:
      label: Steps to reproduce
    validations:
      required: true
  - type: textarea
    attributes:
      label: Expected vs actual
    validations:
      required: true
  - type: input
    attributes:
      label: Docker image tag / stack version
    validations:
      required: true
