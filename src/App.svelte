<script>
	async function getRecipes() {
		let response = await fetch("https://web-carrot.vercel.app/api/getRecipe");
		let recipes = await response.json();
		return recipes;
	}

	async function getFilteredRecipes(filter) {
		let request = "https://web-carrot.vercel.app/api/filter/search?filter=" + filter;
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
	}

	async function addMultipleRecipes(number) {
		for (let index = 0; index < number; index++) {
			await addItem();
		}
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

	import RecipeCard from '../public/card.svelte';
	import { storeFE, undo } from '../scripts/store.js'
	import { fly } from 'svelte/transition'
</script>

<main>
	<div>
		<div style="background:#379683; position:sticky; width:100%; top:0; z-index:1; padding-bottom:10px;">
			<h1>We need to cook.</h1>
			<Container>
				<div style="white-space: nowrap;">
					<Row>
						<Col>
							<Input
								type="text"
								placeholder="Search..."
								bind:value={filter}
								style="width:70%; display:inline-block; margin-left:5px; background:#EDF5E1"
							/>
							<Button style="display: inline-block; margin-left:5px; background:#05386B; color:#EDF5E1" on:click={getFilteredRecipes(filter)} disabled={filter == ""}><Icon name="search" /></Button>
						</Col>
						<Col>
							<Button style="float:right; background:darkred" on:click={removeAllRecipes} disabled={$storeFE.length == 0}>Remove All <Icon name="trash" /></Button>
						</Col>
					</Row>
					<Row>
						<Col>
							<Input
								type="number"
								min={1}
								bind:value
								style="width:40%; display: inline-block; margin-left:5px; background:#EDF5E1"
								placeholder="#"
							/>
							<Button style="display: inline-block; margin-left: 5px; background:#05386B; color:#EDF5E1" on:click={addMultipleRecipes(value)} disabled={value == ""}><Icon name="plus-circle" />{value !== 1 && value !== null ? " Add " + value +" recipes!" : " Add recipe!"}</Button>
						</Col>
						<Col>
							<Button style="float:right; background:#EDF5E1; border:#05386B; color:#05386B" on:click={undoTask} disabled={$undo.length == 0}><Icon name="arrow-counterclockwise" />{$undo.length == 0 ? " Undo" : " Undo " + $undo[$undo.length - 1].task}</Button>
						</Col>
					</Row>
				</div>
			</Container>
		</div>
		<div style="z-index:99">
			{#each $storeFE as recipe}
				<div transition:fly="{{x:-300}}">
					<RecipeCard recipes={recipe}/>
				</div>
			{/each}
		</div>
		<div>
			<img class="gif" src="homer.gif" alt="Gotta do the cooking by the book!"/>
		</div>
	</div>
</main>

<style>
	main {
		text-align: left;
		/* padding: 1em; */
		/* max-width: 240px; */
		/* margin: 0 auto; */
		background-color: #EDF5E1;
		font-family: Georgia, 'Times New Roman', Times, serif
	}

	h1 {
		color: #EDF5E1;
		text-transform: uppercase;
		text-align: center;
		/* font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif */
		/* font-size: 4em; */
		/* font-weight: 100; */
		/* https://visme.co/blog/website-color-schemes/ */
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}

	.gif {
		display: block;
		margin-left: auto;
		margin-right: auto;
		width: 50%;
	}
</style>