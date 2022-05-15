<script>
	async function getRecipes() {
		let response = await fetch("http://localhost:8050/api/getrandom");
		let recipes = await response.json();
		return recipes;
	}

	async function addItem(){
		var l = $storeFE.length;	// get our current items list count
		$storeFE[l] =  await getRecipes();
		console.log($storeFE);
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
		Col
	} from 'sveltestrap';

	import RecipeCard from '../public/card.svelte';
	import { storeFE } from '../scripts/store.js'
</script>

<main>
	<h1>We need to cook.</h1>
	<div>
		<Button size="lg" on:click={addItem}>Add a recipe!</Button>
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