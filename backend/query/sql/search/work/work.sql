SELECT * FROM works
WHERE
    LOWER(
        COALESCE(title_kind, '') || ' ' ||
        COALESCE(title_nickname, '') || ' ' ||
        COALESCE(catalog_prefix, '') || ' ' ||
        COALESCE(catalog_number, '') || ' ' ||
        COALESCE(catalog_subnumber, '')
    ) LIKE '%' || LOWER(?) || '%';
