type IntegrationConvertSuite struct {
    suite.Suite
}

func TestIntegrationConvertSuite(t *testing.T) {
    testSuite := &IntegrationConvertSuite{}
    suite.Run(t, testSuite)

    _ = os.RemoveAll(testdata.ConvertDestination1)
}

func (suite *IntegrationConvertSuite) TestConvertConfigWithEnvironmentsAtTheRoot() {
    transform, err := NewConvert(testdata.ConvertDestination1)
    assert.Nil(suite.T(), err)

    err = transform.ConvertDirectory(testdata.ConvertSource1)
    
    assert.Nil(suite.T(), err)
}