package person

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	clientmocks "github.com/insan1a/exile/internal/client/mocks"
	"github.com/insan1a/exile/internal/models"
	brokermocks "github.com/insan1a/exile/internal/storage/broker/mocks"
	storagemocks "github.com/insan1a/exile/internal/storage/person/mocks"
)

func TestNew(t *testing.T) {
	type args struct {
		options []Option
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "service with nil consumer",
			args:    args{[]Option{WithConsumer(nil)}},
			wantErr: true,
		},
		{
			name:    "service with nil producer",
			args:    args{[]Option{WithProducer(nil, "")}},
			wantErr: true,
		},
		{
			name:    "service with nil storage",
			args:    args{[]Option{WithPeopleStorage(nil)}},
			wantErr: true,
		},
		{
			name:    "service with timeout",
			args:    args{[]Option{WithTimeout(time.Hour)}},
			wantErr: false,
		},
		{
			name:    "service with consumer",
			args:    args{[]Option{WithConsumer(brokermocks.NewConsumer(t))}},
			wantErr: false,
		},
		{
			name:    "service with producer",
			args:    args{[]Option{WithProducer(brokermocks.NewProducer(t), "")}},
			wantErr: false,
		},
		{
			name:    "service with storage",
			args:    args{[]Option{WithPeopleStorage(storagemocks.NewStorage(t))}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := New(tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_Save(t *testing.T) {
	consumer := brokermocks.NewConsumer(t)
	storage := storagemocks.NewStorage(t)
	agify := clientmocks.NewFetcher(t)
	genderize := clientmocks.NewFetcher(t)
	nationalize := clientmocks.NewFetcher(t)
	timeout := time.Second

	svc, err := New(
		WithConsumer(consumer),
		WithPeopleStorage(storage),
		WithTimeout(timeout),
		WithNationalizeClient(nationalize),
		WithAgifyClient(agify),
		WithGenderizeClient(genderize),
	)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	tp := models.Person{
		Name:        "Roman",
		Surname:     "Kravchuk",
		Age:         25,
		Gender:      "male",
		Nationality: "US",
	}

	data, _ := json.Marshal(&tp)

	consumer.On("Consume", timeout).
		Once().
		Return(data, nil)
	agify.On("Fetch", tp.Name).Once().Return([]byte(`{"age":25}`), nil)
	genderize.On("Fetch", tp.Name).Once().Return([]byte(`{"gender":"male"}`), nil)
	nationalize.On("Fetch", tp.Name).Once().Return([]byte(`{"nationality":"US"}`), nil)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	storage.On("Create", ctx, &tp).
		Once().
		Return(nil)

	_, err = svc.Save(ctx)
	if err != nil {
		t.Fatalf("svc.Save() error = %v", err)
	}
}

func TestService_SendErrMessage(t *testing.T) {
	producer := brokermocks.NewProducer(t)

	svc, err := New(
		WithProducer(producer, ""),
	)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	meta := make(map[string]any, 0)
	param := []byte(`{"key": "value"}`)
	_ = json.Unmarshal(param, &meta)

	errStr := "some error"
	msg := models.ErrorMessage{
		Meta:  meta,
		Error: errStr,
	}

	data, _ := json.Marshal(&msg)

	producer.On("Produce", data).
		Once().
		Return(nil)

	err = svc.SendErrMessage(param, errStr)
	if err != nil {
		t.Fatalf("svc.SendErrMessage() error = %v", err)
	}
}
