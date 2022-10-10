

export interface File  {
    size: number,
    format: string,
    href:string,
}

export interface Book {
    id: number|string,
    href: string,
    cover_url: string,
    description?: string,
    author_sort?: string,
    publisher: string,
    pubyear?: string,
    website?: string,
    source?: string,
    provider_key?: string,
    is_owner?: string,
    title?: string,
    author?: string,
    files: File[],
    rating?: number,
    authors?: string[],
    pubdate:string,
    series?: string,
    timestamp?:number,
    isbn?: string,
    collector?: string,
    tags?: string,
}

