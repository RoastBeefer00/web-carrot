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

	import { 
		Button,
		Card,
		CardBody,
		CardFooter,
		CardHeader,
		CardSubtitle,
		CardText,
		CardTitle,
		FormGroup,
		Input,
		ListGroup,
		ListGroupItem,
		Container,
		Row,
		Col,
		Accordion,
		AccordionItem,
		Form
	} from 'sveltestrap';

	let value;
	let filter;

	import RecipeCard from '../public/card.svelte';
	import { storeFE } from '../scripts/store.js'
</script>

<main>
	<h1>We need to cook.</h1>
	<div style="white-space: nowrap">
		<Input
			type="text"
			placeholder="Search for something..."
			bind:value={filter}
			style="width:25%; display: inline-block"
			on:keypress
		/>
		<Button style="display: inline-block; margin-left:20px; background:blue" on:click={getFilteredRecipes(filter)}>Search</Button>
		<p style="display: inline-block; margin-left:20px">or</p>
		<Input
			type="number"
			min={1}
			bind:value
			style="width:25%; display: inline-block; margin-left:20px"
			placeholder="Add # of random recipes..."
		/>
		{#if value!=null}
		<Button style="display: inline-block; margin-left: 20px; background:blue" on:click={addMultipleRecipes(value)}>Add {value} recipe(s)!</Button>
		{/if}
		<Button style="float:right; background:red" on:click={removeAllRecipes}>Remove All</Button>
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
		color: #006115;
		text-transform: uppercase;
		/* font-size: 4em; */
		/* font-weight: 100; */
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>