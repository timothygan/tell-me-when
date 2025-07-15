# Infrastructure

This directory contains infrastructure-as-code for the Tell Me When project.

## Tools
- Terraform
- Terragrunt (for managing environments/modules)

## Suggested Structure
- `/environments` - Per-environment Terragrunt configs (dev, staging, prod)
- `/modules` - Reusable Terraform modules (network, db, app, etc.)
- `/scripts` - Helper scripts for automation

## Getting Started
- See `/environments/README.md` and `/modules/README.md` for details.
