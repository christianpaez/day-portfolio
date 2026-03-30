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

document.querySelectorAll("details").forEach((el) => {
  const content = el.querySelector("p");
const arrow = el.querySelector(".arrow");
  // set initial state
  content.style.overflow = "hidden";
  if (!el.open) content.style.height = "0px";

  el.querySelector("summary").addEventListener("click", (e) => {
    e.preventDefault(); // stop instant toggle

    if (el.open) {
      // CLOSE
      const height = content.scrollHeight;
      arrow.style.transform = "rotate(0deg)";
      content.style.height = height + "px";

      requestAnimationFrame(() => {
        content.style.transition = "height 0.5s ease";
        content.style.height = "0px";
      });

      setTimeout(() => {
        el.open = false;
      }, 300);

    } else {
      // OPEN
      el.open = true;
      arrow.style.transform = "rotate(180deg)";

	    content.style.opacity = "0";
content.style.transform = "translateY(0)";
      const height = content.scrollHeight;
      content.style.height = "0px";

      requestAnimationFrame(() => {
        content.style.transition = "height 0.5s ease";
	      content.style.opacity = "1";
content.style.transform = "translateY(0)";
        content.style.height = height + "px";
      });
    }
  });
});
});

