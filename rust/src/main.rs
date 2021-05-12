use futures::executor::block_on;
use futures::future::join_all;
use reqwest::header::USER_AGENT;
use reqwest::Error;
use serde::Deserialize;
use std::time::SystemTime;

#[derive(Deserialize, Debug)]
struct User {
    login: String,
    id: u32,
}

async fn fetch_users() -> Result<Vec<User>, Error> {
    let request_url = format!(
        "https://api.github.com/repos/{owner}/{repo}/stargazers",
        owner = "rust-lang-nursery",
        repo = "rust-cookbook"
    );
    let client = reqwest::Client::new();
    let response = client
        .get(&request_url)
        .header(USER_AGENT, "My Rust Program 1.0")
        .send()
        .await?;
    let users: Vec<User> = response.json().await?;
    Ok(users)
}

async fn fetch_user(user_struct: &User) -> Result<User, Error> {
    let request_url = format!(
        "https://api.github.com/users/{user}",
        user = user_struct.login
    );
    let now = SystemTime::now();
    println!("fetching user: {:?} @ {:?}", request_url, now);
    let client = reqwest::Client::new();
    let response = client
        .get(&request_url)
        .header(USER_AGENT, "My Rust Program 1.0")
        .send()
        .await?;
        let user: User = response.json().await?;
        Ok(user)
}

#[tokio::main]
async fn main() -> Result<(), Error> {
    match block_on(fetch_users()) {
        Ok(users) => {
            let filtered_users = users
            .iter()
            .filter(|u| u.login.starts_with("s"))
            .map(|u| fetch_user(u));
            let fetch_all_users = join_all(filtered_users).await;
            println!("{:?}", fetch_all_users)
        }
        Err(error) => println!("Oops {:?}", error),
    }
    Ok(())
}