<script>
    export let recipes;

    async function getRecipes() {
		let response = await fetch("http://localhost:8050/api/getrandom");
		let recipes = await response.json();
		return recipes;
	}

    function removeRecipe(title) {
		$storeFE = $storeFE.filter(r => r.title !== title);
	}

    async function replaceRecipe(recipe) {
        for (let index = 0; index < $storeFE.length; index++) {
            if ($storeFE[index] == recipe) {
                $storeFE[index] = await getRecipes();
            } 
        }
    }

    import { storeFE } from '../scripts/store.js'

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
</script>

<div>
    <Card class="mb-3">
        {#await recipes}
            <p>Loading...</p>
        {:then recipes}
            <CardHeader>
                <Button on:click={removeRecipe(recipes.title)}>Remove</Button>
                <Button on:click={replaceRecipe(recipes)}>Replace</Button>
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
        {/await}
    </Card>
</div>