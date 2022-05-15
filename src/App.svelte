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

	async function addMultipleRecipes(number) {
		for (let index = 0; index < number; index++) {
			await addItem();
		}
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

	let value=1;

	import RecipeCard from '../public/card.svelte';
	import { storeFE } from '../scripts/store.js'
</script>

<main>
	<h1>We need to cook.</h1>
	<div>
			<Input
			  type="range"
			  min={1}
			  max={5}
			  bind:value
			  style="width:25%"
			/>
			<Button on:click={addMultipleRecipes(value)}>Add {value} recipe(s)!</Button>
	</div>

	<div>
		<Accordion stayOpen style="width:60%">
			{#each $storeFE as recipe}
				<AccordionItem header="{recipe.title} ({recipe.time})">
					<RecipeCard recipes={recipe}/>
				</AccordionItem>
			{/each}
		</Accordion>
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

	.my-div {
		display: inline
	}
</style>