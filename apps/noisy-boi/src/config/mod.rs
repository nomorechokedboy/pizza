use config::{Environment, File, FileFormat};
use serde::Deserialize;
use std::net::Ipv4Addr;

const CONFIG_PATH: &str = "./src/config/config";

#[derive(Clone, Debug, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Server {
    pub port: u16,
    pub host: bool,
    pub token_secret: String,
}

#[derive(Clone, Debug, Deserialize)]
pub struct Log {
    pub level: String,
}

#[derive(Clone, Debug, Deserialize)]
pub struct Database {
    pub host: String,
    pub name: String,
    pub user: String,
    pub password: String,
    pub port: u16,
}

#[derive(Clone, Debug, Deserialize)]
pub struct AppSettings {
    pub database: Database,
    pub server: Server,
    pub log: Log,
    pub redis: Redis,
}

#[derive(Clone, Debug, Deserialize)]
pub struct Redis {
    pub host: String,
    // pub port: u16,
}

impl AppSettings {
    pub fn new() -> Result<Self, config::ConfigError> {
        let cfg_builder = config::Config::builder()
            .add_source(File::new(CONFIG_PATH, FileFormat::Json))
            .add_source(Environment::default().separator("_"));
        cfg_builder.build()?.try_deserialize::<Self>()
    }

    pub fn server_url(&self) -> String {
        let Server { port, host, .. } = self.server;
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
