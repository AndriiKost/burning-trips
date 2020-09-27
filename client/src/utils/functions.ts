export function stripHtml(html: string): string {
   var tmp = document.createElement("DIV");
   tmp.innerHTML = html;
   return tmp.textContent || tmp.innerText || "";
}

export function jsonCopy(object: any): any {
   return JSON.parse(JSON.stringify(object));
}

export function imageExists(image_url: string) {
   var http = new XMLHttpRequest();
   http.open('HEAD', image_url, false);
   http.send();
   return http.status != 404;
}