use std::str::FromStr;

use crate::utils::parse_generic_command;

use super::ZeroShellCommandsError;

#[derive(Debug, PartialEq, Eq)]
pub struct Pwd {
    pub command: String,
    pub args: Vec<String>,
}

impl Pwd {
    pub fn from_str(command: &str) -> Self {
        let (command, args) = parse_generic_command(command);
        Pwd { command, args }
    }
    pub fn execute(&self) -> Result<(), ZeroShellCommandsError<String>> {
        if self.args.len() > 0 {
            return Err(ZeroShellCommandsError::Pwd(
                "pwd: too many arguments".to_string(),
            ));
        }
        let current_dir = match std::env::current_dir() {
            Ok(dir) => dir,
            Err(error) => {
                eprintln!("Error getting current directory: {}", error);
                return Err(ZeroShellCommandsError::Pwd(error.to_string()));
            }
        };
        println!("{}", current_dir.display());
        Ok(())
    }
}

impl FromStr for Pwd {
    type Err = ();

    fn from_str(command: &str) -> Result<Self, Self::Err> {
        Ok(Pwd::from_str(command))
    }
}

