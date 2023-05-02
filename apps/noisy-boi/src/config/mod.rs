use std::net::Ipv4Addr;

use config::{Environment, File, FileFormat};
use serde::Deserialize;

const CONFIG_PATH: &str = "./src/config/config";

#[derive(Debug, Deserialize)]
pub struct Server {
    pub port: u16,
    pub host: bool,
}

#[derive(Debug, Deserialize)]
pub struct Log {
    pub level: String,
}

#[derive(Debug, Deserialize)]
pub struct Database {
    pub host: String,
    pub name: String,
    pub user: String,
    pub password: String,
    pub port: u16,
}

#[derive(Debug, Deserialize)]
pub struct AppSettings {
    pub database: Database,
    pub server: Server,
    pub log: Log,
    pub redis: Redis,
}

#[derive(Debug, Deserialize)]
pub struct Redis {
    pub host: String,
    pub port: u16,
}

impl AppSettings {
    pub fn new() -> Result<Self, config::ConfigError> {
        let cfg_builder = config::Config::builder()
            .add_source(File::new(CONFIG_PATH, FileFormat::Json))
            .add_source(Environment::default().separator("_"));
        cfg_builder.build()?.try_deserialize::<Self>()
    }

    pub fn server_url(&self) -> String {
        let Server { port, host } = self.server;
        let host = match host {
            true => Ipv4Addr::UNSPECIFIED,
            false => Ipv4Addr::LOCALHOST,
        };
        format!("{host}:{port}")
    }

    pub fn db_url(&self) -> String {
        let Database {
            host,
            name,
            user,
            password,
            port,
        } = &self.database;
        format!("postgres://{user}:{password}@{host}:{port}/{name}")
    }
}
