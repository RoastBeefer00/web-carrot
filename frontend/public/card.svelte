<script>
    async function getRecipes() {
		let response = await fetch("https://mealplanning.azurewebsites.net/api/getrandom");
		let recipes = await response.json();
		return recipes;
	}

    function removeRecipe(recipe) {
        for (let index = 0; index < $storeFE.length; index++) {
            if ($storeFE[index] == recipe) {
                var l = $undo.length;
                $undo[l] = {
                    "task": "Delete",
                    "index": index,
                    "recipe": recipe
                }
            } 
        }
        
        console.log($undo)
		$storeFE = $storeFE.filter(r => r !== recipe);
	}

    async function replaceRecipe(recipe) {
        for (let index = 0; index < $storeFE.length; index++) {
            if ($storeFE[index] == recipe) {
                var l = $undo.length;
                $undo[l] = {
                    "task": "Replace",
                    "index": index,
                    "recipe": recipe
                }
                $storeFE[index] = await getRecipes();
            } 
        }
        console.log($undo)
    }

    function toggle() {
        visible = !visible;
    }

    export let recipes;
    let visible = false;

	import { 
		Button,
		Card,
		CardBody,
		CardHeader,
		CardSubtitle,
		CardTitle,
		Input,
		ListGroup,
		ListGroupItem,
		Container,
		Row,
		Col,
        Icon
	} from 'sveltestrap';

    import { storeFE, undo } from '../scripts/store.js'
    import { slide } from 'svelte/transition';
</script>

<div>
    <Card class="mb-3">
        {#await recipes}
            <p>Loading...</p>
        {:then recipes}
            <CardHeader style="background:#05386B">
                <CardTitle style="color:#EDF5E1;">{recipes.title}</CardTitle>
                <CardSubtitle style="color:#EDF5E1">{recipes.time}</CardSubtitle>
                <Button style="background:#379683; color:#EDF5E1" on:click={toggle}><Icon name={visible ? "chevron-up" : "chevron-down"} /></Button>
                <Button style="float:right; margin-left:20px; background:#379683; color:#EDF5E1" on:click={removeRecipe(recipes)}><Icon name="trash" /></Button>
                <Button style="float:right; margin-left:100px; background:#379683; color:#EDF5E1" on:click={replaceRecipe(recipes)}><Icon name="arrow-repeat" /></Button>
            </CardHeader>
            {#if visible}
            <div transition:slide>
                <CardBody style="background:#8EE4AF">
                    <CardTitle style="color:#05386B">Ingredients</CardTitle>
                    <Container>
                        <Row cols={4}>
                            {#each recipes.ingredients as ingredient}
                                <Col>
                                    <Input style="color:#05386B" id="c1" type ="checkbox" label={ingredient} />
                                </Col>
                            {/each}
                        </Row>
                    </Container>
                    <CardTitle style="color:#05386B">Steps</CardTitle>
                    <ListGroup numbered>
                    {#each recipes.steps as step}
                        <ListGroupItem style="background:#EDF5E1; color:#05386B">{step}</ListGroupItem>
                    {/each}
                    </ListGroup>
                </CardBody>
            </div>
            {/if}
        {/await}
    </Card>
</div>