use serde::{Deserialize, Serialize};
use std::collections::HashMap;

#[derive(Debug, Serialize, Deserialize)]
pub struct Config {
    pub serums: HashMap<String, Serum>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Serum {
    pub source: String,
    pub mask: String,
    pub targets: HashMap<String, Target>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Target {
    pub path: String,
    pub clues: Vec<String>,
}
