package main

func main() {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Add(NewIndividualCustomer("bob"))
	c.Accept(&ServiceRequestVisitor{})

	c2 := &CustomerCol{}
	c2.Add(NewEnterpriseCustomer("A company"))
	c2.Add(NewIndividualCustomer("bob"))
	c2.Add(NewEnterpriseCustomer("B company"))
	c2.Accept(&AnalysisVisitor{})
}
