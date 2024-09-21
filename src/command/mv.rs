use std::{fs, path::Path};

use super::ZeroShellCommandsError;

#[derive(Debug, PartialEq, Eq)]
pub struct Mv {
    pub source: Vec<String>,
    pub destination: Option<String>,
}

impl Mv {
    pub fn from_str(command: &str) -> Self {
        let mut command = command
            .split_whitespace()
            .map(|f| f.to_string())
            .collect::<Vec<String>>();
        if command.len() < 3 {
            return Self {
                source: command[1..command.len()].to_vec(),
                destination: None,
            };
        }
        Self {
            source: command[1..command.len() - 1].to_vec(),
            destination: command.pop(),
        }
    }
    pub fn execute(&self) -> Result<(), ZeroShellCommandsError<String>> {
        dbg!(self.source.clone());
        if self.source.len() == 0 {
            println!("mv: missing file operand");
            return Ok(());
        }
        if self.source.len() > 2 {
            println!(
                "mv: target '{}' is not a directory",
                self.destination.as_ref().unwrap()
            );
            return Ok(());
        }
        if self.destination.is_none() {
            println!(
                "mv: missing destination file operand after '{}'",
                self.source[0]
            );
            return Ok(());
        }

        for source in &self.source {
            let destination = self.destination.as_ref().unwrap().clone();
            if Path::new(&source).exists() {
                if Path::new(&destination).is_dir() {
                    let file_name = match Path::new(&source).file_name() {
                        Some(file_name) => file_name,
                        None => {
                            println!("mv: cannot stat '{}': No such file or directory", source);
                            return Ok(());
                        }
                    };
                    let destination = Path::new(&destination).join(file_name);
                    match fs::rename(source, destination) {
                        Ok(_) => {
                            continue;
                        }
                        Err(err) => {
                            println!("mv: cannot stat '{}': {}", source, err);
                            return Ok(());
                        }
                    }
                }
                match fs::rename(source, destination) {
                    Ok(_) => {
                        continue;
                    }
                    Err(err) => {
                        println!("mv: cannot stat '{}': {}", source, err);
                        return Ok(());
                    }
                }
            } else {
                println!("mv: cannot stat '{}': No such file or directory", source);
            }
        }

        Ok(())
    }
}
