# Product

<!-- impeccable:product-schema 1 -->

## Platform

web

## Users

Independent developers, small engineering teams, and DevOps engineers who need to keep dynamic values synchronized across ordinary files.

## Product Purpose

Devsyringe is a small Go CLI that runs a command, captures a value from its output with a regular expression, and injects that value into one or more files. The landing page succeeds when a visitor understands the mechanism and installs the CLI.

## Positioning

Devsyringe updates plain static files directly without introducing templates, a build system, or runtime configuration hooks.

## Operating Context

Users work in terminals, YAML configuration, local development environments, and CI pipelines. Common inputs include tunnel URLs, tokens, identifiers, and build values.

## Capabilities and Constraints

- A `serum` defines the source command, regex mask, optional timeout, and target files.
- Target lines are located with textual `clues`.
- One captured value can update multiple files.
- The CLI includes a TUI for inspecting processes, statuses, and logs.
- The landing page remains a static HTML/CSS/JavaScript site with no build step.

## Brand Commitments

Keep the Devsyringe name and existing logo. The product voice is concise, technical, and direct.

## Evidence on Hand

- A real animated CLI demonstration at `site/demo2.gif`.
- A working YAML configuration example documented in `README.md`.
- Installation through Go, AUR, and Homebrew.
- The public GitHub repository and a long-form article.
- No testimonials, customer logos, usage metrics, or performance claims are available and none should be fabricated.

## Product Principles

- Explain the run, capture, locate, inject mechanism concretely.
- Make the primary installation command immediate and easy to copy.
- Let real configuration and demo material earn trust.
- Keep secondary paths, including GitHub and the article, discoverable without competing with installation.
