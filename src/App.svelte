<script>
	async function getRecipes() {
		let response = await fetch("http://localhost:8050/api/getrandom");
		let recipes = await response.json();
		return recipes;
	}
	// const promise = getRecipes();

	// let recipes = async function  getRecipes() {
    // 	let response = await fetch("http://localhost:8050/api/getrandom");
	// 	console.log(response.json)
	// 	return response.json
    //   }

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

	let visible = false
	function toggle(){
		visible = !visible
	}

	import RecipeCard from '../public/card.svelte';
</script>

<main>
	{#if !visible}
	<div>
		<Button size="lg" on:click={toggle}>I am a button</Button>
	</div>
	{/if}
	<!-- <Recipes/> -->
	<div>
		{#if visible}
		<Button on:click={toggle}> button 2 </Button>
		<RecipeCard recipes={getRecipes()}/>
		<RecipeCard recipes={getRecipes()}/>
		<RecipeCard recipes={getRecipes()}/>
		<RecipeCard recipes={getRecipes()}/>

			

		<!-- <Card class="mb-3">
			{#await promise}
				<p>Loading...</p>
			{:then recipes}
				<CardHeader>
					<CardTitle>{recipes.title}</CardTitle>
					<CardSubtitle>Prep Time: {recipes.time}</CardSubtitle>
				</CardHeader>
				<CardBody>
					<CardTitle>Ingredients</CardTitle>
					<Container>
						<Row cols={4}>
							{#each recipes.ingredients as ingredient}
								<Col>
									<Input id="c1" type ="checkbox" label={ingredient} />
								</Col>
							{/each}
						</Row>
					</Container>
					<h1>Steps</h1>
					<ListGroup numbered>
					{#each recipes.steps as step}
						<ListGroupItem>{step}</ListGroupItem>
					{/each}
					</ListGroup>
				</CardBody>
			{:catch error}
				<p style="color: red">{error.message}</p>
			{/await}
		</Card> -->
		{/if}
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