SELECT Orders.Id, Orders.Name, Orders.StreetAddr, Orders.City, 
    Orders.Shipped
FROM Orders
WHERE Orders.Id = ?