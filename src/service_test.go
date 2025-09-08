package test

type ServiceTestSuite struct {
	suite.Suite

	svc Service
}

func newConfigService() fflags.ConfigService {
	return idgenmock.NewMockConfigService(fakeflags.NewFakeConfigService().WithInt32Value("my-int-feature-flag", 1))
}

func (s *ServiceTestSuite) newTestService() Service {
	// BEGIN __INCLUDE_DB__
	db := newDB(s.T())
	// END __INCLUDE_DB__
	// BEGIN __INCLUDE_LAUNCHDARKLY__
	configService := newConfigService()
	// BEGIN __INCLUDE_DB__
	widgetIDGenerator, err := mockWidgetIDGenerator(db, configService)
	require.NoError(s.T(), err)
	// END __INCLUDE_DB__
	// END __INCLUDE_LAUNCHDARKLY__
	return &Impl{
		Clock:  newClock(),
		Params: newParams(),
		// BEGIN __INCLUDE_LAUNCHDARKLY__
		ConfigService: configService,
		// END __INCLUDE_LAUNCHDARKLY__
		MeterProvider: metricnoop.NewMeterProvider(),
		// BEGIN __INCLUDE_DB__
		DB: db,
		// BEGIN __INCLUDE_LAUNCHDARKLY__
		WidgetIDGenerator: widgetIDGenerator,
		// END __INCLUDE_LAUNCHDARKLY__
		// END __INCLUDE_DB__
	}
}
func (s *ServiceTestSuite) SetupSuite() {
	s.svc = s.newTestService()
}
func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
func (s *Impl) SayHello(
	ctx context.Context, request *api.SayHelloRequest,
) (*api.SayHelloResponse, error) {
	// Examples of logging and metrics. APIs invoked are captured in logging middleware.
	// Logging in individual endpoint is not required.
	ctxlog.Info(ctx, "hello endpoint called")

	// the metric will look like io.confluent.cc-go-template-service.say_hello_count
	counter, err := s.MeterProvider.
		Meter("io.confluent.cc-go-template-service").
		Int64Counter("say_hello_count")
	if err != nil {
		return nil, err
	}
	counter.Add(ctx, 1)

	// BEGIN __INCLUDE_LAUNCHDARKLY__
	val, err := s.ConfigService.Int32Value(ctx, "my-int-feature-flag", nil)
	if err != nil {
		return nil, err
	}
	// END __INCLUDE_LAUNCHDARKLY__
	ctxlog.Debug(ctx, "something interesting happened")
	message := "hello " + request.Name +
		// BEGIN __INCLUDE_LAUNCHDARKLY__
		" - you get " + strconv.Itoa(int(val)) +
		// END __INCLUDE_LAUNCHDARKLY__
		" - Foo is " + s.Params.Foo
	return &api.SayHelloResponse{Message: message, DateTime: s.Clock.Now()}, nil
}
