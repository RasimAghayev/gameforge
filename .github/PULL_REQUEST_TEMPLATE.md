name: Pull request template
description: Template for pull requests
title: "[feat] "
body:
  - type: textarea
    attributes:
      label: Description
    validations:
      required: true
  - type: input
    attributes:
      label: Related issue
  - type: dropdown
    attributes:
      label: Change type
      options: [feat, fix, refactor, perf, test, build, ci, docs, chore, revert]
    validations:
      required: true
  - type: input
    attributes:
      label: Verification command + exit code
    validations:
      required: true
  - type: checkboxes
    attributes:
      label: Checklist
      options:
        - label: Tests updated
        - label: Docs updated
        - label: Risk / rollback noted
        - label: No AI attribution in commit/PR
    validations:
      required: true
