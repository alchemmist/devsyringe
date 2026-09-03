---
name: Devsyringe
description: A flat technical-document system for a small developer tool.
colors:
  ink: "#1c2940"
  paper: "#f6f8f5"
  link: "#4055a8"
  command-surface: "#e6f0dc"
  config-surface: "#e8eef7"
  divider: "#cbd6d8"
  muted: "#526071"
typography:
  display:
    fontFamily: "ui-monospace, SFMono-Regular, Cascadia Code, Roboto Mono, Consolas, monospace"
    fontSize: "clamp(36px, 7vw, 52px)"
    fontWeight: 700
    lineHeight: 1.04
    letterSpacing: "-0.055em"
  body:
    fontFamily: "ui-monospace, SFMono-Regular, Cascadia Code, Roboto Mono, Consolas, monospace"
    fontSize: "15px"
    fontWeight: 400
    lineHeight: 1.75
  code:
    fontFamily: "ui-monospace, SFMono-Regular, Cascadia Code, Roboto Mono, Consolas, monospace"
    fontSize: "14px"
    fontWeight: 400
    lineHeight: 1.7
spacing:
  page-gutter: "20px"
  section: "56px"
components:
  code-surface:
    backgroundColor: "{colors.command-surface}"
    textColor: "{colors.ink}"
    typography: "{typography.code}"
    padding: "19px 21px"
---

# Design System: Devsyringe

## Overview

**Creative North Star: "The Project README"**

The site behaves like a concise, carefully typeset technical document. It uses one readable column, a system monospace face, real code, and explicit blue links. Product evidence is presented directly rather than packaged inside promotional UI.

## Colors

Blue-graphite ink sits on a cool laboratory-paper field. Indigo is reserved for actions and destinations. Install commands use a pale reagent green; configuration uses a cool instrument blue.

## Typography

One system monospace stack serves every role. Hierarchy comes from weight, size, spacing, and line length rather than switching typefaces. Body copy stays at 15px with 1.75 line height; commands stay at 14px or larger.

## Layout

All content follows a single 680px column with 20px minimum side gutters. Sections stack in document order and use 56px vertical padding. Structured values may use compact two-column definition rows, collapsing naturally when needed.

## Elevation & Depth

The system is flat and uses no shadows. Light gray code surfaces and one-pixel rules provide separation.

## Shapes

Corners are square. The only shapes are the existing product logo and browser-native text underlines.

## Components

### Code surfaces

Commands sit on pale reagent green and YAML sits on cool instrument blue, both with 14px monospace text and direct underlined copy actions.

### Links

Links use indigo `#4055a8`, a three-pixel underline offset, and a thicker underline on hover.

### Definition rows

Labels are bold and muted; values use full-contrast text. One-pixel dividers preserve scanability.

## Do's and Don'ts

### Do:

- **Do** keep the page as one readable technical document.
- **Do** use real commands, configuration, and demonstrations as the visual content.
- **Do** keep body and command text at readable sizes with strong contrast.

### Don't:

- **Don't** introduce marketing-card grids, oversized empty hero regions, shadows, gradients, glass, or decorative technical motifs.
- **Don't** load a display font when the system monospace stack communicates the product more honestly.
- **Don't** use pale gray for information the visitor needs to understand or act on.
