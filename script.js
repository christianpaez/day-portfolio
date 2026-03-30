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
  const summary = el.querySelector("summary");

  // initial styles
  content.style.overflow = "hidden";
  content.style.transition = "height 0.5s ease, opacity 0.3s ease, transform 0.3s ease";

  if (!el.open) {
    content.style.height = "0px";
    content.style.opacity = "0";
    content.style.transform = "translateY(-6px)";
    arrow.style.transform = "rotate(0deg)";
  }

  function closeAccordion() {
    const height = content.scrollHeight;
    content.style.height = height + "px";

    content.offsetHeight; // force repaint

    content.style.height = "0px";
    content.style.opacity = "0";
    content.style.transform = "translateY(-6px)";
    arrow.style.transform = "rotate(0deg)";

    content.addEventListener("transitionend", function handler(e) {
      if (e.propertyName !== "height") return;
      el.open = false;
      content.removeEventListener("transitionend", handler);
    });
  }

  function openAccordion() {
    el.open = true;

    const height = content.scrollHeight;

    content.style.height = "0px";
    content.style.opacity = "0";
    content.style.transform = "translateY(-6px)";

    content.offsetHeight; // force repaint

    content.style.height = height + "px";
    content.style.opacity = "1";
    content.style.transform = "translateY(0)";
    arrow.style.transform = "rotate(180deg)";
  }

  summary.addEventListener("click", (e) => {
    e.preventDefault();
    el.open ? closeAccordion() : openAccordion();
  });

  // ✅ click content to close
  content.addEventListener("click", () => {
    if (!el.open) return;
    closeAccordion();
  });
});
});

