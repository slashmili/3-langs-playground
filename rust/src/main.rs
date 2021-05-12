use futures::executor::block_on;
use reqwest::header::USER_AGENT;
use reqwest::Error;
use serde::Deserialize;

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
    println!("{}", request_url);
    let client = reqwest::Client::new();
    let response = client
        .get(&request_url)
        .header(USER_AGENT, "My Rust Program 1.0")
        .send()
        .await?;
    let users: Vec<User> = response.json().await?;
    return Ok(users);
}

#[tokio::main]
async fn main() -> Result<(), Error> {
    match block_on(fetch_users()) {
        Ok(users) => {
            let filtered_users = users.iter().filter(|u| u.login.starts_with("s"));
            println!("{:?}", filtered_users);
        }
        Err(error) => println!("Oops {:?}", error),
    }

    /*
     */
    Ok(())
}
