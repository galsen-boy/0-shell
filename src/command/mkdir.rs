use std::fs;

use crate::utils::parse_generic_command;

use super::ZeroShellCommandsError;

#[derive(Debug, PartialEq, Eq)]
pub struct Mkdir {
    pub file: Vec<String>,
}

impl Mkdir {
    pub fn from_str(command: &str) -> Self {
        let (_command, file) = parse_generic_command(command);
        Self { file }
    }
    pub fn execute(&self) -> Result<(), ZeroShellCommandsError<String>> {
        for file in &self.file {
            match fs::create_dir(file) {
                Ok(_) => {}
                Err(e) => {
                    println!("{}", e);
                    return Ok(());
                }
            }
        }
        Ok(())
    }
}
