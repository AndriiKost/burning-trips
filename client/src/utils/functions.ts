export function stripHtml(html: string): string {
   var tmp = document.createElement("DIV");
   tmp.innerHTML = html;
   return tmp.textContent || tmp.innerText || "";
}

export function jsonCopy(object: any): any {
   return JSON.parse(JSON.stringify(object));
}