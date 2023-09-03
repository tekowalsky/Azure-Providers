# Azure-Providers

### Main branch status
![main branch](https://github.com/tekowalsky/azure-providers/actions/workflows/getSchemas.yml/badge.svg?branch=main)

### Dev branch status
![dev branch](https://github.com/tekowalsky/azure-providers/actions/workflows/getSchemas.yml/badge.svg?branch=dev)

## Purpose

This is a small convenience project to check Terraform provider support for currently available Azure resource types.  This workflow collects Azure resource provider and resource type lists from the Microsoft Azure ARM schema repository and matches them to Hashicorp Azurerm provider schemas from Hashicorp repositories.  

## Output

The schemas and other output files are committed to this repository's "schemas" branch twice each day.
