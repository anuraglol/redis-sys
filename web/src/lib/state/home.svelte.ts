export class HomeState {
	selectedKey = $state("");
	formKey = $state("");
	formValue = $state("");
	formTtl = $state("");

	resetForm() {
		this.formKey = "";
		this.formValue = "";
		this.formTtl = "";
	}
}

export const homeState = new HomeState();
