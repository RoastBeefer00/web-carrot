<script>
	async function getRecipes() {
		let response = await fetch("http://localhost:8050/api/getrandom");
		let recipes = await response.json();
		return recipes;
	}

	async function getFilteredRecipes(filter) {
		let request = "http://localhost:8050/api/filter/" + filter;
		console.log(request);
		let response = await fetch(request);
		let recipes = await response.json();

		var l = $storeFE.length;
		 recipes.recipes.forEach(element => {
			$storeFE[l] = element;
			l ++;
		});	
	}

	async function addItem(){
		var l = $storeFE.length;	// get our current items list count
		$storeFE[l] =  await getRecipes();
		console.log($storeFE);
	}

	async function addMultipleRecipes(number) {
		for (let index = 0; index < number; index++) {
			await addItem();
		}
	}

	function removeAllRecipes() {
		$storeFE = [];
	}

	let value;
	let filter;

	import { 
		Button,
		Input,
		Icon
	} from 'sveltestrap';

	import RecipeCard from '../public/card.svelte';
	import { storeFE } from '../scripts/store.js'
</script>

<main>
	<div style="background:#379683">
	<h1>We need to cook.</h1>
	<div style="white-space: nowrap;">
		<Input
			type="text"
			placeholder="Search for something..."
			bind:value={filter}
			style="width:25%; display: inline-block; margin-left:20px; background:#EDF5E1"
			on:keypress
		/>
		<Button style="display: inline-block; margin-left:20px; background:#05386B; color:#EDF5E1" on:click={getFilteredRecipes(filter)}><Icon name="search" /></Button>
		<p style="display: inline-block; margin-left:20px; color:white">or</p>
		<Input
			type="number"
			min={1}
			bind:value
			style="width:25%; display: inline-block; margin-left:20px; background:#EDF5E1"
			placeholder="Add # of random recipes..."
		/>
		<Button style="display: inline-block; margin-left: 20px; background:#05386B; color:#EDF5E1" on:click={addMultipleRecipes(value)}><Icon name="plus-circle" />{value !== 1 && value !== null ? " Add " + value +" recipes!" : " Add recipe!"}</Button>
		<Button style="float:right; background:darkred; margin-right:20px" on:click={removeAllRecipes}>Remove All <Icon name="trash" /></Button>
	</div>
	</div>

	<div>
		{#each $storeFE as recipe}
			<RecipeCard recipes={recipe}/>
		{/each}
	</div>
</main>

<style>
	main {
		text-align: left;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #EDF5E1;
		text-transform: uppercase;
		text-align: center;
		/* font-size: 4em; */
		/* font-weight: 100; */
		/* https://visme.co/blog/website-color-schemes/ */
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>