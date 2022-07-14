package services

var CREATE_PRODUCT = `INSERT INTO public.product (
	name, description, price, specification, category_id, created_at, updated_at)
	VALUES ('%s', '%s', %f, '%s', %d, '%s', '%s');`

var DELETE_PRODUCT = `DELETE FROM public.product WHERE public.product.id=%d`

var UPDATE_PRODUCT = `UPDATE public.product 
	SET name='%s', description='%s', price=%f, specification='%s', category_id=%d, updated_at='%s', rating=%d
	WHERE public.product.id=%d`

var SELECT_PRODUCT = `SELECT * FROM public.product WHERE public.product.id=%d`
var SELECT_ALL_PRODUCT = `SELECT * FROM public.product`
