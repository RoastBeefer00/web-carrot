<script>
	async function getRecipes(num) {
		let response = await fetch("https://hae0pt.deta.dev/random?num="+num);
		let recipes = await response.json();
		var l = $storeFE.length;
		recipes.forEach(element => {
			$storeFE[l] = element;
			l ++;
		});	
	}

	async function getFilteredRecipes(filter) {
		let request = "https://hae0pt.deta.dev/recipes?filter=" + filter;
		console.log(request);
		let response = await fetch(request);
		let recipes = await response.json();

		var l = $storeFE.length;
		recipes.forEach(element => {
			$storeFE[l] = element;
			l ++;
		});	
	}

	function removeAllRecipes() {
		$storeFE = [];
	}

	async function undoTask() {
		var undoLength = $undo.length;
		var recipesLength = $storeFE.length;
		var item = $undo[undoLength - 1];
		if (item.task == "Delete") {
			for (let index = ($storeFE.length - 1); index >= item.index; index--) {
				$storeFE[index + 1] = await $storeFE[index];
				console.log($storeFE);
			}
			$storeFE[item.index] = item.recipe;
		}
		else {
			$storeFE[item.index] = item.recipe;
		}

		$undo = $undo.filter(r => r.recipe !== item.recipe);
	}

	let value;
	let filter;

	import { 
		Button,
		Input,
		Icon,
		Col,
		Container,
		Row
	} from 'sveltestrap';

	import { invoke } from '@tauri-apps/api/tauri'

	async function getRandomRecipe(num) {
		let recipe = await invoke('get_random_recipe');
		console.log(recipe);
		// let recipe =  await response.json();
		recipe = JSON.parse(recipe);
		$storeFE[$storeFE.length] = recipe;
	}

	import RecipeCard from '../public/card.svelte';
	import { storeFE, undo } from '../scripts/store.js'
	import GroceryList  from '../public/GroceryList.svelte';
	import { fly } from 'svelte/transition'
</script>

<main>
	<div style="height:100%">
		<div class="header">
			<h1>We need to cook.</h1>
			<Container>
				<div style="white-space: nowrap;">
					<Row>
						<Col>
							<input
								type="text"
								placeholder="Search..."
								bind:value={filter}
								class="form search_bar"
							/>
							<button 
								class="button button_normal" 
								on:click={getFilteredRecipes(filter)} 
								disabled={filter == ""}>
								<Icon name="search" />
							</button>
						</Col>
						<Col>
							<button
								class="button button_delete button_right"
								on:click={removeAllRecipes} 
								disabled={$storeFE.length == 0}>
								Remove All <Icon name="trash" />
						</button>
						</Col>
					</Row>
					<Row>
						<Col>
							<input
								type="number"
								min={1}
								bind:value
								class="form add_recipes_form"
								placeholder="#"
							/>
							<button 
								class="button button_normal" 
								on:click={getRandomRecipe} 
								disabled={value == undefined}>
								<Icon name="plus-circle" />
								{value !== 1 && value !== undefined ? " Add " + value +" recipes!" : " Add recipe!"}
							</button>
						</Col>
						<Col>
							<button 
								class="button button_normal button_right" 
								on:click={undoTask} 
								disabled={$undo.length == 0}>
								<Icon name="arrow-counterclockwise" />
								{$undo.length == 0 ? " Undo" : " Undo " + $undo[$undo.length - 1].task}
							</button>
						</Col>
					</Row>
				</div>
			</Container>
			<div style="text-align:center;">
				<GroceryList />
			</div>	
		</div>
		<div class="recipe_store">
			{#each $storeFE as recipe}
				<div class="recipe" transition:fly="{{x:-300}}">
					<RecipeCard recipes={recipe}/>
				</div>
			{/each}
		</div>
		<div>
			<img 
				class="gif" 
				src="homer.gif" 
				alt="Gotta do the cooking by the book!"
			/>
		</div>
	</div>
</main>

<style>
	main {
		text-align: left;
		background-color: #EDF5E1;
		font-family: Georgia, 'Times New Roman', Times, serif;
	}

	/* HEADER */

	h1 {
		color: #EDF5E1;
		text-transform: uppercase;
		text-align: center;
	}

	.header {
		background: #379683; 
		position: sticky; 
		width: 100%; 
		top: 0; 
		z-index: 1; 
		padding-bottom: 10px;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}

	/* BUTTONS */

	.button_normal {
		background-color: #05386B; 
		color: #EDF5E1;
	}

	.button_delete {
		background-color: darkred;
		color: #EDF5E1;
	}

	.button:disabled {
		background-color: darkslategrey;
	}


	/* FORMS */

	.form {
		display: inline-block; 
		margin: 2px; 
		background: #EDF5E1;
		border-radius: 8px;
	}

	.add_recipes_form {
		width: 40%; 
		
	}

	.search_bar {
		width: 70%; 
	}

	/* RECIPES */
	.recipe {
		margin-bottom: 10px;
	}

	/* BOTTOM HALF */

	.recipe_store {
		z-index: 99;
		height: 100%; 
		flex-grow: 1;
		padding: 20px;
	}

	.gif {
		display: block;
		margin-left: auto;
		margin-right: auto;
		width: 50%;
	}
</style>
