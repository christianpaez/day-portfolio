document.addEventListener("DOMContentLoaded", (event) => {
	const headerMenuButton = document.querySelector("#header-menu-button");
	const headerMenu = document.querySelector("#header-menu");
	const openHeaderMenuIcon = document.querySelector("#open-header-menu-icon");
	const closeHeaderMenuIcon = document.querySelector("#close-header-menu-icon");

	headerMenuButton.addEventListener("click", (event) => {
		headerMenu.classList.toggle("hidden");
		openHeaderMenuIcon.classList.toggle("hidden");
		closeHeaderMenuIcon.classList.toggle("hidden");
	});
});

