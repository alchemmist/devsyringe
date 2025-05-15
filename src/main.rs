mod config_model;

use config_model::Config;
use std::fs::File;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let file = File::open("./injection.yaml")?;
    let config: Config = serde_yaml::from_reader(file)?;
    println!("{:#?}", config.serums["https_host"]);
    Ok(())
}
