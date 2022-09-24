#![cfg_attr(
  all(not(debug_assertions), target_os = "windows"),
  windows_subsystem = "windows"
)]

use serde::{Deserialize, Serialize};

use std::path::Path;
use std::error::Error;
use std::fs::File;
use std::io::BufReader;
use rand::Rng;

#[derive(Serialize, Deserialize, Debug)]
struct Root {
    recipes: Vec<Recipe>,
}

#[derive(Serialize, Deserialize, Debug)]
struct Recipe {
    title: String,
    time: String,
    ingredients: Vec<String>,
    steps: Vec<String>,
}

fn read_user_from_file<P: AsRef<Path>>(path: P) -> Result<Root, Box<dyn Error>> {
    // Open the file in read-only mode with buffer.
    let file = File::open(path)?;
    let reader = BufReader::new(file);

    // Read the JSON contents of the file as an instance of `Root`.
    let u = serde_json::from_reader(reader)?;

    // Return the `Root`.
    Ok(u)
}

#[tauri::command]
fn get_random_recipe() -> String {
    let u = read_user_from_file("recipes.json").unwrap();
    let recipes = u.recipes;
    let mut rng = rand::thread_rng();
    let length = recipes.len();
    let rand = rng.gen_range(0..length);

    let ret = serde_json::to_string(&recipes[rand]).unwrap();
    return ret
}

fn main() {
  tauri::Builder::default()
    .invoke_handler(tauri::generate_handler![get_random_recipe])
    .run(tauri::generate_context!())
    .expect("error while running tauri application");
}
