SELECT transactions_detail.id, transactions.order_at, transactions.paid_at, transactions.status_paid, transactions_detail.qty, transactions_detail.sub_total
FROM transactions_detail
INNER JOIN transactions ON transactions_detail.transaction_id = transactions.id