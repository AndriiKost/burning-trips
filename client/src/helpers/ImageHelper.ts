export function constructImageUrlForWikipedia(imageThumbUrl: string): string {
    const imageWithoutThumb = imageThumbUrl.replace('/thumb/', '/').split('/');
    imageWithoutThumb.pop();
    return imageWithoutThumb.join('/');
}

// https://upload.wikimedia.org/wikipedia/commons/thumb/8/8e/Camp_Randall_Stadium_aerial_%28cropped%29.jpg/200px-Camp_Randall_Stadium_aerial_%28cropped%29.jpg
// https://upload.wikimedia.org/wikipedia/commons/8/8e/Camp_Randall_Stadium_aerial_%28cropped%29.jpg